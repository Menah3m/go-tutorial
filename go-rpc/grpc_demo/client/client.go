package main

import (
	"context"
	"fmt"
	"github.com/menah3m/go-tutorial/go-rpc/grpc_demo/proto"
	"google.golang.org/grpc"
)

/*
   @Auth: menah3m
   @Desc:
*/

func main() {
	// stream 模式
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic("failed to connect:" + err.Error())

	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "kangkang", Gender: 0, Mp: map[string]string{
		"name":    "kangmp",
		"company": "keymap",
	}})

	if err != nil {
		panic("failed to call:" + err.Error())
	}
	fmt.Println(r.Message, r.Gender.String())
	r1, _ := c.Ping(context.Background(), &proto.Empty{})
	fmt.Println(r1.Id)
}
