package main

import (
	"context"
	"log"
	"time"

	pb "gRPC/server/proto"

	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewHelloClient(conn)

	// Make a request to the server
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Ishan"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Client received: %s", response.GetMessage())
}
