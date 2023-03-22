package main

import (
	"context"
	pb "github.com/menah3m/go-tutorial/grpc-demo/helloworld"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

/*
   @Auth: menah3m
   @Desc: grpc client
*/

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect : %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer c
	cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{
		Name: name,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)

}
