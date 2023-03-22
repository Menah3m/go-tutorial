package main

import (
	"fmt"
	"net/rpc"
)

/*
   @Auth: menah3m
   @Desc:
*/

func main() {
	// 1 建立连接
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	var reply string
	err = client.Call("HelloService.Hello", "kangminjie", &reply)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)

}
