package main

import (
	"context"
	"flag"
	"github.com/be9/dual-grpc-go/proto/testservice"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	// Define command line flag for the hostname
	host := flag.String("host", "", "hostname to connect to")
	flag.Parse()

	// Create a gRPC connection to the specified host
	conn, err := grpc.Dial(*host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client
	client := testservice.NewTestServiceClient(conn)

	// Make a gRPC request
	response, err := client.GetFoos(context.Background(), &testservice.GetFoosRequest{})
	if err != nil {
		if statusErr, ok := status.FromError(err); ok {
			log.Fatalf("gRPC error: %v (%d)", statusErr.Message(), statusErr.Code())
		} else {
			log.Fatalf("failed to make gRPC request: %v", err)
		}
	}

	log.Printf("gRPC response: %v", response)
}
