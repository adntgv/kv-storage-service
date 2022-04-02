package main

import (
	"context"
	"log"

	pb "github.com/adntgv/kv-storage-service/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type server struct {
	db *gorm.DB
	pb.UnimplementedKeyValueServer
}

type Pair struct {
	gorm.Model
	Key   string
	Value string
}

type Event struct {
	gorm.Model
	Event string
	Key   string
	Value string
}

func newPair(in *pb.Pair) *Pair {
	return &Pair{
		Key:   in.Key,
		Value: in.Value,
	}
}

func (pair *Pair) toPair() *pb.Pair {
	return &pb.Pair{
		Key:   pair.Key,
		Value: pair.Value,
	}
}

/*

If a user tries to create an answer that already exists -
the request should fail and an adequate message or code should be returned.
*/
func (s *server) Create(ctx context.Context, in *pb.Pair) (*pb.Reply, error) {
	if result := s.db.First(new(Pair)).Where("key", in.Key); result.Error == nil {
		return nil, status.Errorf(codes.AlreadyExists, "already exists")
	} else if result.Error != gorm.ErrRecordNotFound {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}

	pair := newPair(in)

	if result := s.db.Create(pair); result.Error != nil {
		return nil, status.Errorf(codes.InvalidArgument, result.Error.Error())
	}

	go s.record("create", pair)

	return success(), nil
}

func (s *server) record(event string, data *Pair) {
	res := s.db.Create(&Event{Event: event, Key: data.Key, Value: data.Value})
	if res.Error != nil {
		log.Println("could not save event", res.Error.Error())
	}
}

func success() *pb.Reply {
	return &pb.Reply{Response: &pb.Reply_Message{Message: "success"}}
}

/*

If a user saves the same key multiple times (using update),
every answer should be saved. When retrieving an answer,
it should return the latest answer.

it is not possible to update a deleted key.
*/
func (s *server) Update(ctx context.Context, in *pb.Pair) (*pb.Reply, error) {
	if result := s.db.First(new(Pair)).Where("key", in.Key); result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.Internal, result.Error.Error())
		}
		return nil, status.Errorf(codes.NotFound, "not found")
	}

	pair := newPair(in)

	if result := s.db.Create(pair); result.Error != nil {
		return nil, status.Errorf(codes.InvalidArgument, result.Error.Error())
	}

	go s.record("update", pair)

	return success(), nil
}

/*

// returns the latest answer for the given key

If an answer doesn't exist or has been deleted,
an adequate message or code should be returned.
*/
func (s *server) Get(ctx context.Context, in *pb.Key) (*pb.Reply, error) {
	pair := new(Pair)
	if result := s.db.Last(pair).Where("key", in.Key); result.Error != nil {
		return nil, status.Errorf(codes.InvalidArgument, result.Error.Error())
	}

	return &pb.Reply{Response: &pb.Reply_Pair{Pair: pair.toPair()}}, nil
}

func (s *server) Delete(ctx context.Context, in *pb.Key) (*pb.Reply, error) {
	pair := new(Pair)
	if result := s.db.Delete(pair).Where("key", in.Key); result.Error != nil {
		return nil, status.Errorf(codes.InvalidArgument, result.Error.Error())
	}

	pair.Key = in.Key
	go s.record("delete", pair)

	return success(), nil
}

/*
//returns an array of events in chronological order

When returning history, only mutating events (create, update, delete) should be returned.
The "get" events should not be recorded.
*/
func (s *server) GetHistory(ctx context.Context, in *pb.Key) (*pb.HistoryReply, error) {
	events := make([]*Event, 0)
	if result := s.db.Find(&events).Where("key", in.Key); result.Error != nil {
		return nil, status.Errorf(codes.InvalidArgument, result.Error.Error())
	}

	historyEvents := make([]*pb.Event, len(events))

	for i := 0; i < len(events); i++ {
		historyEvents[i] = new(pb.Event)
		historyEvents[i].Data = make(map[string]string)

		historyEvents[i].Event = events[i].Event
		historyEvents[i].Data[events[i].Key] = events[i].Value
	}

	return &pb.HistoryReply{Events: historyEvents}, nil
}

func (s *server) Clear(ctx context.Context, in *pb.Key) (*pb.Reply, error) {
	if result := s.db.Where("1 = 1").Delete(&Pair{}); result.Error != nil {
		return nil, status.Errorf(codes.InvalidArgument, result.Error.Error())
	}

	if result := s.db.Where("1 = 1").Delete(&Event{}); result.Error != nil {
		return nil, status.Errorf(codes.InvalidArgument, result.Error.Error())
	}

	return success(), nil
}
