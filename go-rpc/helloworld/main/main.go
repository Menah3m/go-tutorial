package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	helloworld "github.com/menah3m/go-tutorial/go-rpc/helloworld/proto"
)

/*
   @Auth: menah3m
   @Desc:
*/

type Hello struct {
	Name string `json:"name"`
}

func main() {
	req := helloworld.HelloRequest{Name: "kangkang"}
	rsp, _ := proto.Marshal(&req)
	jsonStruct := Hello{Name: "kangkang"}
	jsonRsp, _ := json.Marshal(jsonStruct)
	fmt.Println(string(jsonRsp))
	fmt.Println(string(rsp))
}
