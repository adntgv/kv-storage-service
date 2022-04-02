package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/adntgv/kv-storage-service/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:5001", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewKeyValueClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Create(ctx, &pb.Pair{Key: "name", Value: "Aidyn"})
	if err != nil {
		log.Fatalf("could not create: %v", err)
	}
	log.Printf("Answer: %s", r.GetMessage())

	r, err = c.Get(ctx, &pb.Key{Key: "name"})
	if err != nil {
		log.Fatalf("could not create: %v", err)
	}
	log.Printf("Answer: %v", r.String())
}
