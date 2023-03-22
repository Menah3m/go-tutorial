package main

import (
	"fmt"
	"io/ioutil"
)

/*
   @Auth: menah3m
   @Desc:
*/

func main() {
	buf, err := ioutil.ReadFile("./1.txt")
	if err != nil {
		fmt.Println(err)
	}
	s := string(buf)
	fmt.Printf("%s", s)

}
