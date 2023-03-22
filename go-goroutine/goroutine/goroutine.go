package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
   @Auth: menah3m
   @Desc:
*/

func printGoroutineNo(i int) {
	for {
		fmt.Printf("hello from goroutine %d\n", i)
	}
}

func multiPrint() {
	for i := 0; i < 10; i++ {
		go printGoroutineNo(i)
	}
	time.Sleep(30 * time.Second)
}

func go1(ch chan string) {
	for i := 0; i < 20; i++ {
		ch <- "I'm goroutine1. Num:" + strconv.Itoa(i)
		time.Sleep(5 * time.Second)
	}
}

func go2(ch chan int) {
	for i := 0; i < 20; i++ {
		ch <- i
		time.Sleep(60 * time.Second)
	}

}

func multiChan() {
	ch1 := make(chan string, 3)
	ch2 := make(chan int, 5)

	for i := 0; i < 10; i++ {
		go go1(ch1)
		go go2(ch2)
	}

	for {
		select {
		case str, ch1check := <-ch1:
			if !ch1check {
				fmt.Println("ch1 failed.")
			}
			fmt.Println(str)
		case p, ch2check := <-ch2:
			if !ch2check {
				fmt.Println("ch2 failed.")
			}
			fmt.Println(p)
		}

	}
}

func main() {
	multiPrint()
	// multiChan()
}
