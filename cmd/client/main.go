package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/adntgv/kv-storage-service/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr  = flag.String("addr", "localhost:5001", "the address to connect to")
	key   = flag.String("key", "name", "")
	value = flag.String("value", "Madonna", "")
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

	cmds := map[string]func(pb.KeyValueClient) error{
		"test":       test,
		"create":     Create,
		"update":     Update,
		"get":        Get,
		"delete":     Delete,
		"getHistory": GetHistory,
		"clear":      Clear,
	}

	if err := cmds[os.Args[1]](c); err != nil {
		log.Fatalf("%v %v", os.Args[1], err)
	}
}

func Create(c pb.KeyValueClient) error {
	in := &pb.Pair{
		Key:   *key,
		Value: *value,
	}
	res, err := c.Create(context.Background(), in)
	if err != nil {
		return fmt.Errorf("Create: %v", err)
	}

	log.Printf("%v", res.GetResponse())

	return nil
}
func Update(c pb.KeyValueClient) error {
	in := &pb.Pair{
		Key:   *key,
		Value: *value,
	}
	res, err := c.Update(context.Background(), in)
	if err != nil {
		return fmt.Errorf("Update: %v", err)
	}

	log.Printf("%v", res.GetResponse())

	return nil
}
func Get(c pb.KeyValueClient) error {
	in := &pb.Key{
		Key: *key,
	}
	res, err := c.Get(context.Background(), in)
	if err != nil {
		return fmt.Errorf("Get: %v", err)
	}

	log.Printf("%v", res.GetResponse())

	return nil
}
func Delete(c pb.KeyValueClient) error {
	in := &pb.Key{
		Key: *key,
	}
	res, err := c.Delete(context.Background(), in)
	if err != nil {
		return fmt.Errorf("Delete: %v", err)
	}

	log.Printf("%v", res.GetResponse())

	return nil
}
func GetHistory(c pb.KeyValueClient) error {
	in := &pb.Key{
		Key: *key,
	}
	res, err := c.GetHistory(context.Background(), in)
	if err != nil {
		return fmt.Errorf("GetHistory: %v", err)
	}

	log.Printf("%v", res.GetEvents())

	return nil
}
func Clear(c pb.KeyValueClient) error {
	in := &pb.Key{}
	res, err := c.Clear(context.Background(), in)
	if err != nil {
		return fmt.Errorf("Clear: %v", err)
	}

	log.Printf("%v", res.GetResponse())

	return nil
}

func test(c pb.KeyValueClient) error {
	for i, f := range []func(pb.KeyValueClient) error{
		createExisting,
		updateAfterDelete,
		history,
	} {
		if err := f(c); err != nil {
			return fmt.Errorf("function %v did not pass: %v", i, err)
		}
	}

	return nil
}

func createExisting(c pb.KeyValueClient) error {
	log.Println("Testing double creation")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Clear(ctx, &pb.Key{})
	if err != nil {
		return fmt.Errorf("could not create: %v", err)
	}
	log.Printf("Answer: %s", r.GetMessage())

	_, err = c.Create(ctx, &pb.Pair{Key: "name", Value: "Aidyn"})
	if err != nil {
		return fmt.Errorf("could not create: %v", err)
	}

	_, err = c.Create(ctx, &pb.Pair{Key: "name", Value: "Aidyn"})
	if err == nil {
		return fmt.Errorf("should not be able to create: %v", err)
	}
	log.Printf("Answer: %s", err.Error())

	return nil
}

func updateAfterDelete(c pb.KeyValueClient) error {
	log.Println("Testing update deleted")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	log.Println("clear db")
	r, err := c.Clear(ctx, &pb.Key{})
	if err != nil {
		return fmt.Errorf("could not create: %v", err)
	}
	log.Printf("Answer: %s", r.GetMessage())

	log.Println("create pair")
	r, err = c.Create(ctx, &pb.Pair{Key: "name", Value: "Aidyn"})
	if err != nil {
		return fmt.Errorf("could not create: %v", err)
	}
	log.Printf("Answer: %s", r.GetMessage())

	log.Println("delete pair")
	r, err = c.Delete(ctx, &pb.Key{Key: "name"})
	if err != nil {
		return fmt.Errorf("could not delete: %v", err)
	}
	log.Printf("Answer: %s", r.GetMessage())

	log.Println("update deleted")
	_, err = c.Update(ctx, &pb.Pair{Key: "name", Value: "Aidyn"})
	if err == nil {
		return fmt.Errorf("should not be able to update: %v", err)
	}
	log.Printf("Answer: %s", err.Error())

	return nil
}

func history(c pb.KeyValueClient) error {
	log.Println("Testing get history")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	log.Println("clear db")
	r, err := c.Clear(ctx, &pb.Key{})
	if err != nil {
		return fmt.Errorf("could not create: %v", err)
	}
	log.Printf("Answer: %s", r.GetMessage())

	log.Println("create pair")
	r, err = c.Create(ctx, &pb.Pair{Key: "name", Value: "Aidyn"})
	if err != nil {
		return fmt.Errorf("could not create: %v", err)
	}
	log.Printf("Answer: %s", r.GetMessage())

	log.Println("update")
	r, err = c.Update(ctx, &pb.Pair{Key: "name", Value: "Aidyn"})
	if err != nil {
		return fmt.Errorf("should not be able to update: %v", err)
	}
	log.Printf("Answer: %s", r.GetMessage())

	log.Println("delete pair")
	r, err = c.Delete(ctx, &pb.Key{Key: "name"})
	if err != nil {
		return fmt.Errorf("could not delete: %v", err)
	}
	log.Printf("Answer: %s", r.GetMessage())

	log.Println("get history pair")
	hist, err := c.GetHistory(ctx, &pb.Key{Key: "name"})
	if err != nil {
		return fmt.Errorf("could not delete: %v", err)
	}
	log.Printf("Answer: %v", hist.GetEvents())

	return nil
}
