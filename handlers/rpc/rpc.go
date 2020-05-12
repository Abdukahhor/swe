package rpc

import (
	"context"

	"github.com/abdukahhor/swe/app"
	pb "github.com/abdukahhor/swe/handlers/rpcpb"
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

//GetNumber - Возвращает число по id инкремента. В самом начале это ноль.
func (s *server) GetNumber(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	reply := s.core.Get(in.Id)
	return &pb.Response{Number: reply.Num, Code: reply.Code, Message: reply.Message, Id: in.Id}, nil
}

//IncrementNumber - Увеличивает число на размер инкремента.
func (s *server) IncrementNumber(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	reply := s.core.Increment(in.Id)
	return &pb.Response{Code: reply.Code, Message: reply.Message, Id: in.Id}, nil
}

//SetSettings - настройка инкремента, size это размер инкремента, max это размер верхней границы
//возвращает id инкремента, и по этой id делаеться инкремент (IncrementNumber) и возвращает число (GetNumber)
func (s *server) SetSettings(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	reply := s.core.Set(in.Id, in.Size, in.Max)
	return &pb.Response{Code: reply.Code, Message: reply.Message, Id: reply.ID}, nil
}
