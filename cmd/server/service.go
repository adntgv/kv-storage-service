package main

import (
	"context"

	pb "github.com/adntgv/kv-storage-service/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type server struct {
	db *gorm.DB
	pb.UnimplementedKeyValueServer
}

func (s *server) Create(ctx context.Context, in *pb.Pair) (*pb.Reply, error) {
	if result := s.db.Create(in); result.Error != nil {
		return nil, status.Errorf(codes.InvalidArgument, result.Error.Error())
	}

	return &pb.Reply{Response: &pb.Reply_Message{Message: "success"}}, nil
}

func (s *server) Update(ctx context.Context, in *pb.Pair) (*pb.Reply, error) {
	if result := s.db.Create(in); result.Error != nil {
		return nil, status.Errorf(codes.InvalidArgument, result.Error.Error())
	}

	return &pb.Reply{Response: &pb.Reply_Message{Message: "success"}}, nil
}

func (s *server) Get(ctx context.Context, in *pb.Key) (*pb.Reply, error) {
	pair := &pb.Pair{}
	if result := s.db.Last(pair).Where("name", in.Key); result.Error != nil {
		return nil, status.Errorf(codes.InvalidArgument, result.Error.Error())
	}

	return &pb.Reply{Response: &pb.Reply_Pair{Pair: pair}}, nil
}

func (s *server) Delete(ctx context.Context, in *pb.Key) (*pb.Reply, error) {
	if result := s.db.Delete(in); result.Error != nil {
		return nil, status.Errorf(codes.InvalidArgument, result.Error.Error())
	}

	return &pb.Reply{Response: &pb.Reply_Message{Message: "success"}}, nil
}

func (s *server) GetHistory(ctx context.Context, in *pb.Key) (*pb.HistoryReply, error) {
	if result := s.db.Find(in); result.Error != nil {
		return nil, status.Errorf(codes.InvalidArgument, result.Error.Error())
	}

	events := make([]*pb.Event, 0)

	return &pb.HistoryReply{Events: events}, nil
}
