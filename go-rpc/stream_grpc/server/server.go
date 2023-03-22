package main

import (
	"fmt"
	"github.com/menah3m/go-tutorial/go-rpc/stream_grpc/proto"
	"google.golang.org/grpc"
	"net"
	"sync"
	"time"
)

/*
   @Auth: menah3m
   @Desc:
*/
const PORT = ":9090"

type Server struct {
}

func (e *Server) GetStream(req *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	i := 0
	for true {
		i++
		_ = res.Send(&proto.StreamResData{Data: fmt.Sprintf("%s %v", req.Data, time.Now().Unix())})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}

	return nil

}

func (e *Server) PutStream(cliStr proto.Greeter_PutStreamServer) error {

	for true {
		a, err := cliStr.Recv()
		if err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(a.Data)
		}
	}

	return nil

}

func (e *Server) All(allStr proto.Greeter_AllServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for true {
			data, _ := allStr.Recv()
			fmt.Println("receive client message:" + data.Data)
		}
	}()

	go func() {
		defer wg.Done()
		for true {
			allStr.Send(&proto.StreamResData{Data: "this message is from server"})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	return nil
}

func main() {
	g := grpc.NewServer()
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}
	proto.RegisterGreeterServer(g, &Server{})
	err = g.Serve(lis)
	if err != nil {
		panic(err)
	}
}
