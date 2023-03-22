package main

import (
	"context"
	"fmt"
	"github.com/menah3m/go-tutorial/go-rpc/stream_grpc/proto"
	"google.golang.org/grpc"
	"sync"
	"time"
)

/*
   @Auth: menah3m
   @Desc:
*/

func main() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	// 服务端 stream 模式
	res, _ := c.GetStream(context.Background(), &proto.StreamReqData{
		Data: "get server stream:",
	})

	for true {

		a, err := res.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println(a.Data)
	}

	// 客户端 stream 模式
	putStr, _ := c.PutStream(context.Background())
	i := 0
	for true {
		i++
		putStr.Send(&proto.StreamReqData{Data: fmt.Sprintf("timestamp %v", time.Now().Unix())})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}

	// 双向 stream 模式
	allStr, _ := c.All(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for true {
			data, _ := allStr.Recv()
			fmt.Println("receive server message:" + data.Data)
		}
	}()

	go func() {
		defer wg.Done()
		for true {
			allStr.Send(&proto.StreamReqData{Data: "this message is from client"})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
}
