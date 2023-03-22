package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
)

/*
   @Auth: menah3m
   @Desc:
*/

type ResponseData struct {
	Data int `json:"data"`
}

func Add(a, b int) int {

	req := HttpRequest.NewRequest()
	url := fmt.Sprintf("http://127.0.0.1:8080/add?a=%d&b=%d", a, b)
	// fmt.Println(url)
	res, _ := req.Get(url)
	body, _ := res.Body()
	r := ResponseData{}
	_ = json.Unmarshal(body, &r)
	return r.Data
}

func main() {

	fmt.Println(Add(3, 4))
}
