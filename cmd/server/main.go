package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/adntgv/kv-storage-service/gen"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedKeyValueServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) Create(ctx context.Context, in *pb.Pair) (*pb.Reply, error) {
	log.Printf("Received: %v", in.GetKey())
	return &pb.Reply{Message: "Hello "}, nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterKeyValueServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
