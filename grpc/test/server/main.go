package main

import (
	"context"
	gproto "go-example/grpc/proto"
	"net"

	"google.golang.org/grpc"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, request *gproto.HelloRequest) (*gproto.HelloReply, error) {
	return &gproto.HelloReply{
		Message: "Hello " + request.Name,
	}, nil
}

func main() {
	g := grpc.NewServer()
	s := Server{}
	gproto.RegisterGreeterServer(g, &s)
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	_ = g.Serve(listen)
}
