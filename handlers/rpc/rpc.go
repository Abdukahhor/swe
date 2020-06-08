package rpc

import (
	"context"

	"github.com/abdukahhor/swe/app"
	pb "github.com/abdukahhor/swe/handlers/rpcpb"
	"github.com/abdukahhor/swe/models"
	"google.golang.org/grpc"
)

//server
type server struct {
	core *app.Core
	// todo log
}

//Register RPC Increment Server and grpc Server
func Register(s *grpc.Server, c *app.Core) {
	pb.RegisterIncrementServer(s, &server{core: c})
}

//GetNumber - Returns a number by the increment id.
func (s *server) GetNumber(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	reply := s.core.Get(in.Id)
	return &pb.Response{Number: reply.Num, Code: reply.Code, Message: reply.Message, Id: in.Id}, nil
}

//IncrementNumber - Increments the number by increment size.
func (s *server) IncrementNumber(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	reply := s.core.Increment(in.Id)
	return &pb.Response{Code: reply.Code, Message: reply.Message, Id: in.Id}, nil
}

//SetSettings - increment setting, size is the size of the increment, max is the size of the upper border
//returns the id of the increment, and this id is used to increment (IncrementNumber) and returns a number (GetNumber)
func (s *server) SetSettings(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	reply := s.core.Set(models.Settings{ID: in.Id, Size: in.Size, Max: in.Max})
	return &pb.Response{Code: reply.Code, Message: reply.Message, Id: reply.ID}, nil
}
