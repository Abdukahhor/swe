package api

import (
	"context"

	"github.com/abdukahhor/swe/app"
	"google.golang.org/grpc"
)

//server
type server struct {
	c *app.Core
	// todo log
}

//Register RPC Increment Server and grpc Server
func Register(s *grpc.Server, p *app.Core) {
	RegisterIncrementServer(s, &server{c: p})
}

//GetNumber -
func (s *server) GetNumber(ctx context.Context, in *Request) (*Response, error) {
	return nil, nil
}

//IncrementNumber -
func (s *server) IncrementNumber(ctx context.Context, in *Request) (*Response, error) {
	return nil, nil
}

//SetSettings -
func (s *server) SetSettings(ctx context.Context, in *Request) (*Response, error) {
	return nil, nil
}
