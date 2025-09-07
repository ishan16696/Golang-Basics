package main

import (
	"context"
	"fmt"
	pb "gRPC/server/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

// Server implements the Greeter service defined in the .proto file.
type server struct {
	pb.UnimplementedHelloServer
}

// SayHello implements the SayHello RPC method.
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	message := fmt.Sprintf("Hello Hii, %s!", req.GetName())
	return &pb.HelloResponse{Message: message}, nil
}

func main() {
	// Set up a TCP listener on port 9000
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterHelloServer(grpcServer, &server{})

	log.Println("gRPC server is running on port 9000")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("gRPC server failed to serve: %v at port: 9000", err)
	}
}
