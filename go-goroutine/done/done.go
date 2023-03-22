package main

import (
	"fmt"
	"sync"
)

/*
   @Auth: menah3m
   @Desc:
*/
func main() {
	chanDemo()

}

type worker struct {
	in chan int
	wg *sync.WaitGroup
}

func chanDemo() {
	var workers [10]worker
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	for i, w := range workers {

		w.in <- 'A' + i
		wg.Add(1)
	}
	// for _, w := range workers {
	// 	<-w.done
	// }

	for i, w := range workers {
		w.in <- 'a' + i
		wg.Add(1)
	}
	wg.Wait()
	//
	// for _, w := range workers {
	// 	<-w.done
	// }

}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int, 3),
		wg: wg,
	}

	go doWork(w.in, id, wg)

	return w
}

func doWork(c chan int, id int, wg *sync.WaitGroup) {
	for n := range c {
		fmt.Printf("worker %d received %c \n", id, n)
		wg.Done()
	}
}
