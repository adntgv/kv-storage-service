package main

import (
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/adntgv/kv-storage-service/gen"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=user password=password dbname=db port=5432 sslmode=disable TimeZone=Asia/Almaty"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to open connection: %v", err)
	}

	db.AutoMigrate(&pb.Pair{})

	port := getEnv("port", "5001")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterKeyValueServer(s, &server{db: db})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func getEnv(key string, defaultValue string) string {
	if val := os.Getenv(key); val == "" {
		return defaultValue
	} else {
		return val
	}
}
