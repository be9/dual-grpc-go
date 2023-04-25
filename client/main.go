package main

import (
	"context"
	"flag"
	"log"
	"math/rand"
	"time"

	"github.com/be9/dual-grpc-go/proto/testservice"
	grpcnew "github.com/be9/grpc-go"
	insecure2 "github.com/be9/grpc-go/credentials/insecure"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	host := flag.String("host", "", "hostname to connect to")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	conn, err := grpc.Dial(*host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client
	client := testservice.NewTestServiceClient(conn)

	conn2, err := grpcnew.Dial(*host, grpcnew.WithTransportCredentials(insecure2.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn2.Close()

	// Create a gRPC client
	client2 := testservice.NewTestServiceClient2(conn2)

	dual := testservice.NewDualClient(client, client2, func() bool {
		return rand.Intn(2) == 1
	})

	for n := 0; n < 10; n++ {
		// Make a gRPC request
		response, err := dual.GetFoos(context.Background(), &testservice.GetFoosRequest{})
		if err != nil {
			if statusErr, ok := status.FromError(err); ok {
				log.Fatalf("gRPC error: %v (%d)", statusErr.Message(), statusErr.Code())
			} else {
				log.Fatalf("failed to make gRPC request: %v", err)
			}
		}

		log.Printf("gRPC response: %v", response)
	}
}
