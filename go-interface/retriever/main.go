package main

import (
	"fmt"
	real2 "github.com/menah3m/go-tutorial/go-interface/retriever/real"
)

/*
   @Auth: menah3m
   @Desc:
*/

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("https://www.imooc.com")
}

func main() {
	var r Retriever
	// r := mock.Retriever{Contents: "this is a fake imooc.com"}
	r = real2.Retriever{Url: "https://www.imooc.com"}
	getTypeOfRetriever(r)
	realRetriever, ok := r.(real2.Retriever)
	if ok {
		fmt.Println(realRetriever.Url)
	} else {
		fmt.Println("not real retriever")
	}
}

func getTypeOfRetriever(r Retriever) {
	switch v := r.(type) {
	case real2.Retriever:
		fmt.Println("Url:", v.Url)
	default:
		fmt.Println("unknown")
	}
}
