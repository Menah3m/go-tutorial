package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
   @Auth: menah3m
   @Desc:
*/

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func createWorker(id int) chan int {
	c := make(chan int)
	go worker(c, id)
	return c
}

func worker(c chan int, id int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("worker %d received %d \n", id, n)

	}
}

func main() {
	c1 := generator()
	c2 := generator()
	var w = createWorker(0)
	var values []int
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)

	for {
		var activeWorker chan int
		var activeValue int
		if len(values) > 0 {
			activeWorker = w
			activeValue = values[0]
		}

		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-tick:
			fmt.Println("values len = ", len(values))
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tm:
			fmt.Println("Byebye.")
			return
		}
	}

}
