package rpc

import (
	"context"
	"net"

	"github.com/abdukahhor/swe/app"
	pb "github.com/abdukahhor/swe/handlers/rpcpb"
	"github.com/abdukahhor/swe/models"
	"google.golang.org/grpc"
)

var defaultServer *grpc.Server

//server
type server struct {
	core *app.Core
	// todo log
}

//Register RPC Increment Server and grpc Server
func Register(c *app.Core) {
	defaultServer = grpc.NewServer()
	pb.RegisterIncrementServer(defaultServer, &server{core: c})
}

//Run serve grpc Server
func Run(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer l.Close()
	err = defaultServer.Serve(l)
	if err != nil {
		return err
	}
	return nil
}

//Close stop grpc Server
func Close() {
	defaultServer.GracefulStop()
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
