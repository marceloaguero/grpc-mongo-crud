package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	blogpb "../proto"
)

// BlogServiceServer implements BlogServiceServer interface
type BlogServiceServer struct{}

type BlogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}

func (s *BlogServiceServer) CreateBlog(ctx context.Context, req *blogpb.CreateBlogReq) (*blogpb.CreateBlogRes, error) {
	// Essentially doing req.Blog to access the struct with a nil check
	blog := req.GetBlog()

	// Now we have to convert this into a BlogItem type to convert to BSON
	data := BlogItem{
		// ID:	Empty, so it gets omitted and mongodb generates a unique Object ID upon insertion
		AuthorID: blog.GetAuthorId(),
		Title:    blog.GetTitle(),
		Content:  blog.GetContent(),
	}

	// Insert the data into de database, result contains the newly generated Object ID for the new document
	result, err := blogdb.InsertOne(mongoCtx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	// Add the id blog, first cast the "generic type" (go doesn't have real generics yes)  to an Object ID
	oid := result.InsertedID.(primitive.ObjectID)

	// Convert the object id to it's string counterpart
	blog.Id = oid.Hex()

	// Return the blog in a CreateBlogRes type
	return &blogpb.CreateBlogRes{Blog: blog}, nil
}

func (s *BlogServiceServer) ReadBlog(ctx context.Context, req *blogpb.ReadBlobReq) (*blogpb.ReadBlogRes, error) {
	// Convert string id (from proto) to mongodb ObjectId
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	result := blogdb.FindOne(ctx, bson.M{"_id": oid})

	// Create an empty BlogItem to write our decode result to
	data := BlogItem{}

	// Decode and write to data
	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find blog with ObjectId %s: %v", req.GetId(), err))
	}

	// Cast to ReadBlogRes
	response := &blogpb.ReadBlogRes{
		Blog: &blogpb.Blog{
			Id:       oid.Hex(),
			AuthorId: data.AuthorID,
			Title:    data.Title,
			Content:  data.Content,
		},
	}
	return response, nil
}

func (s *BlogServiceServer) UpdateBlog(ctx context.Context, req *blogpb.UpdateBlogReq) (*blogpb.UpdateBlogRes, error) {
	// Get the blog data from the request
	blog := req.GetBlog()

	// Convert the Id string to a mongodb ObjectId
	oid, err := primitive.ObjectIDFromHex(blog.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Could not convert the supplied id to a mongodb ObjectId: %v", err),
		)
	}

	// Convert the data to be updated into an unordered bson document
	update := bson.M{
		"author_id": blog.GetAuthorId(),
		"title":     blog.GetTitle(),
		"content":   blog.GetContent(),
	}

	// Convert the oid into an unordered bson to search in by id
	filter := bson.M{"_id": oid}

	// Result is the bson encoded result
	// To return the updated document instead of original, we have to add options
	result := blogdb.FindOneAndUpdate(ctx, filter, bson.M{"$set": update}, options.FindOneAndUpdate().SetReturnDocument(1))

	// Decode result and write it to 'decoded'
	decoded := BlogItem{}
	err = result.Decode(&decoded)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Could not find blog with supplied ID: %v", err),
		)
	}
	return &blogpb.UpdateBlogRes{
		Blog: &blogpb.Blog{
			Id:       decoded.ID.Hex(),
			AuthorId: decoded.AuthorID,
			Title:    decoded.Title,
			Content:  decoded.Content,
		},
	}, nil
}

func (s *BlogServiceServer) DeleteBlog(ctx context.Context, req *blogpb.DeleteBlogReq) (*blogpb.DeleteBlogRes, error) {
	// Get the ID (string) from de request message and convert it to an ObjectId
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}

	// DeleteOne returns DeleteResult which is a struct containing the amount of deleted docs (in this case only 1 always)
	// So we return a boolean instead
	_, err = blogdb.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find/delete blog with id %s: %v", req.GetId(), err))
	}

	// Return response with success: true if no error is thrown (and thus document is removed)
	return &blogpb.DeleteBlogRes{
		Success: true,
	}, nil
}

func (s *BlogServiceServer) ListBlogs(req *blogpb.ListBlogReq, stream blogpb.BlogService_ListBlogsServer) error {
	// Initiate a BlogItem type to write decoded data to
	data := &BlogItem{}

	// Collection.Find returns a cursor for our (empty) query
	cursor, err := blogdb.Find(context.Background(), bson.M{})
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}

	// An expression with defer will be called at the end of the function
	defer cursor.Close(context.Background())

	// cursor.Next() returns a boolean, if false thereare no more items and loop will breake
	for cursor.Next(context.Background()) {
		// Decode the data at the current pointer and write it to data
		err = cursor.Decode(data)
		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}

		// Send blog over stream
		stream.Send(&blogpb.ListBlogRes{
			Blog: &blogpb.Blog{
				Id:       data.ID.Hex(),
				AuthorId: data.AuthorID,
				Title:    data.Title,
				Content:  data.Content,
			},
		})
	}
	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown cursor error: %v", err))
	}
	return nil
}

// Global variables for db connection , collection and context
var db *mongo.Client
var blogdb *mongo.Collection
var mongoCtx context.Context

func main() {
	// Configure log package to give file name and line number on eg. log.Fatal
	// just the filename & line number:
	// log.SetFlags(log.Lshortfile)
	// Or add timestamps and pipe file name line number to it:
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Starting server on port :50051...")

	// 50051 is the default port for gRPC
	// Ideally we'd user 0.0.0.0 instead of localhost as well
	listener, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Unable to listen on port :50051: %v", err)
	}

	// Set a slice of gRPC options
	// Here we can configure things like TLS
	opts := []grpc.ServerOption{}

	// Create a new gRPC server with (blank) options
	// var s *grpc.Server
	s := grpc.NewServer(opts...)

	// Create BlogService type
	// var srv *BlogServiceServer
	srv := &BlogServiceServer{}

	// Register the service with de server
	blogpb.RegisterBlogServiceServer(s, srv)

	// Initialize MongoDb client
	fmt.Println("Connecting to mongodb")

	// non-nil empty context
	mongoCtx := context.Background()

	// Connect takes in a context and options,
	// the connection URI is the only option we pass for now
	db, err := mongo.Connect(mongoCtx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	// Check whether the connection was successful by pinging the mongodb server
	err = db.Ping(mongoCtx, nil)
	if err != nil {
		log.Fatal("Could not connect to mongodb: %v\n", err)
	} else {
		fmt.Println("Connected to mongodb")
	}

	// Bind our collection to our global variable for use in other methods
	blogdb = db.Database("mydb").Collection("blog")

	// Start the server in a child routine
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	fmt.Println("Server successfully started on port :50051")

	// Wrong way to stop de server
	// if err := s.Serve(listener); err != nil {
	// 	log.Fatalf("Failed to serve: %v", err)
	// }

	// Right way to stop the server using a SHUTDOWN HOOK
	// Create a channel to receive OS signals
	c := make(chan os.Signal)

	// Relay os.Interrupt to our channel (os.Interrupt = CTRL+C)
	// Ignore other incoming signals
	signal.Notify(c, os.Interrupt)

	// Block main routine until a signal is received
	// As long as user doesn't press CTRL+C a message is not passed
	// and our main routine keeps running
	<-c

	// Alter receiving CTRL+C, properly stop the server
	fmt.Println("\nStopping the server...")
	s.Stop()
	listener.Close()
	fmt.Println("Closing mongodb connection")
	db.Disconnect(mongoCtx)
	fmt.Println("Done.")
}
