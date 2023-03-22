package main

import (
	"fmt"
	"github.com/menah3m/go-tutorial/go-basic/infra"
	"github.com/menah3m/go-tutorial/go-basic/testing"
)

/*
   @Auth: menah3m
   @Desc:
*/

func getTestingRetriever() retriever {
	return testing.Retriever{}
}

func getInfraRetriever() retriever {
	return infra.Retriever{}
}

type retriever interface {
	Get(string) string
}

func main() {
	retriever := getInfraRetriever()

	bytes := retriever.Get("https://www.imooc.com")
	fmt.Printf("%s \n", bytes)
}
