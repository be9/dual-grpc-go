package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/be9/dual-grpc-go/server/util"
	"golang.org/x/net/trace"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
	"net/http"

	pb "github.com/be9/dual-grpc-go/proto/testservice"
	"google.golang.org/grpc"
)

var (
	port     = flag.Int("port", 50051, "The server port")
	httpPort = flag.Int("http-port", 50052, "The HTTP port")
)

type server struct {
	pb.UnimplementedTestServiceServer
}

func genDetails(n int) []string {
	strings := make([]string, n)
	for i := range strings {
		strings[i] = util.RandStringBytesMask(50)
	}
	return strings
}

func getUserAgent(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "unknown"
	}
	if agent, ok := md["user-agent"]; ok {
		return agent[0]
	}
	return "unknown"
}

func (s *server) GetFoos(ctx context.Context, req *pb.GetFoosRequest) (*pb.GetFoosResponse, error) {
	log.Printf("GetFoos [agent %v]", getUserAgent(ctx))
	foos := make([]*pb.Foo, 10)
	for i := range foos {
		foos[i] = &pb.Foo{
			Id:      fmt.Sprintf("%d", i),
			Details: genDetails(5),
		}
	}
	return &pb.GetFoosResponse{Foos: foos}, nil
}

func (s *server) GetFoobars(ctx context.Context, req *pb.GetFoobarsRequest) (*pb.GetFoobarsResponse, error) {
	log.Printf("GetFoobars [foo_id=%s,agent=%v]", req.GetFooId(), getUserAgent(ctx))
	foobars := make([]*pb.Foobar, 10)
	for i := range foobars {
		foobars[i] = &pb.Foobar{
			Id:      fmt.Sprintf("%d", i),
			Details: genDetails(10),
		}
	}
	return &pb.GetFoobarsResponse{Foobars: foobars}, nil
}

func main() {
	flag.Parse()

	trace.AuthRequest = func(req *http.Request) (bool, bool) {
		return true, true
	}
	grpc.EnableTracing = true

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTestServiceServer(s, &server{})
	log.Printf("gRPC server listening at %v", lis.Addr())

	go (func() {
		log.Printf("Starting HTTP at :%d", *httpPort)
		_ = http.ListenAndServe(fmt.Sprintf(":%d", *httpPort), nil)
	})()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
