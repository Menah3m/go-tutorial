package main

import (
	"fmt"
	"time"
)

/*
   @Auth: menah3m
   @Desc:
*/

func main() {
	// chanDemo()
	// bufferedChannel()
	channelClose()

}

func chanDemo() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
		go worker(channels[i], i)
	}
	for i := 0; i < 10; i++ {

		channels[i] <- 'A' + i
	}
	time.Sleep(10 * time.Second)

}

func createWorker(id int) chan int {
	c := make(chan int)
	go worker(c, id)
	return c
}

func worker(c chan int, id int) {
	for n := range c {
		fmt.Printf("worker %d received %c \n", id, n)

	}
}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(c, 0)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int, 3)
	go worker(c, 0)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}
