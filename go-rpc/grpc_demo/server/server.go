package main

import (
	"context"
	"github.com/menah3m/go-tutorial/go-rpc/grpc_demo/proto"
	"google.golang.org/grpc"
	"net"
)

/*
   @Auth: menah3m
   @Desc:
*/

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: "hello," + request.Name, Gender: request.Gender}, nil
}

func (s *Server) Ping(ctx context.Context, request *proto.Empty) (*proto.Pong, error) {
	return &proto.Pong{Id: "kangminjie"}, nil
}

func main() {
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic("failed to listen:" + err.Error())

	}
	err = g.Serve(lis)
	if err != nil {
		panic("failed to start:" + err.Error())
	}
}
