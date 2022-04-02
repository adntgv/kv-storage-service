package main

import (
	"context"
	"flag"
	"fmt"
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

	for i, f := range []func(pb.KeyValueClient) error{
		create,
		createExisting,
		update,
		updateAfterDelete,
		getDeleted,
		history,
	} {
		if err := f(c); err != nil {
			log.Fatalf("function %v did not pass: %v", i, err)
		}
	}
}

func create(c pb.KeyValueClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Clear(ctx, &pb.Key{})
	if err != nil {
		return fmt.Errorf("could not create: %v", err)
	}
	log.Printf("Answer: %s", r.GetMessage())

	r, err = c.Create(ctx, &pb.Pair{Key: "name", Value: "Aidyn"})
	if err != nil {
		return fmt.Errorf("could not create: %v", err)
	}
	log.Printf("Answer: %s", r.GetMessage())

	return nil
}

func createExisting(c pb.KeyValueClient) error {
	return fmt.Errorf("unimplemented")
}
func update(c pb.KeyValueClient) error {
	return fmt.Errorf("unimplemented")
}
func updateAfterDelete(c pb.KeyValueClient) error {
	return fmt.Errorf("unimplemented")
}
func getDeleted(c pb.KeyValueClient) error {
	return fmt.Errorf("unimplemented")
}
func history(c pb.KeyValueClient) error {
	return fmt.Errorf("unimplemented")
}
