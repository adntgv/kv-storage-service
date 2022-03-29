package main

import (
	"context"

	pb "github.com/adntgv/kv-storage-service/gen"
	"gorm.io/gorm"
)

type server struct {
	db *gorm.DB
	pb.UnimplementedKeyValueServer
}

func (s *server) Create(ctx context.Context, in *pb.Pair) (*pb.Reply, error) {
	if result := s.db.Create(in); result.Error != nil {
		return nil, result.Error
	}

	return &pb.Reply{Message: "Successfully created pair"}, nil
}
