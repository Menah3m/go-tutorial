package main

import (
	"net"
	"net/rpc"
)

/*
   @Auth: menah3m
   @Desc:

*/

type HelloService struct {
}

func (h *HelloService) Hello(request string, reply *string) error {
	*reply = "hello," + request
	return nil
}

func main() {
	// 1 实例化一个 server
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		return
	}
	// 2 注册 handler

	err = rpc.RegisterName("HelloService", &HelloService{})
	if err != nil {
		return
	}

	// 3 启动服务
	for {
		conn, _ := listener.Accept()
		go rpc.ServeConn(conn)
	}

}
