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

func msgGen(name string, done chan struct{}) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			select {
			case <-time.After(time.Duration(rand.Intn(2000)) * time.Millisecond):
				c <- fmt.Sprintf("service: %s  get message:%d", name, i)
			case <-done:
				fmt.Println("cleaning up.")
				time.Sleep(2 * time.Second)
				fmt.Println("cleanup done.")
				done <- struct{}{}
				return
			}
			i++
		}
	}()

	return c
}

func fanIn(c1, c2 chan string) chan string {
	c := make(chan string)

	go func() {
		for {
			c <- <-c1
		}
	}()

	go func() {
		for {
			c <- <-c2
		}
	}()

	return c
}

func fanIns(chs ...chan string) chan string {
	c := make(chan string)
	for _, ch := range chs {

		go func(chCopy chan string) {
			for {
				c <- <-chCopy
			}
		}(ch)
	}
	return c
}

func fanInBySelect(c1, c2 chan string) chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case m := <-c1:
				c <- m
			case m := <-c2:
				c <- m
			}

		}
	}()
	return c
}

func nonBlockingWait(c chan string) (string, bool) {
	select {
	case m := <-c:
		return m, true
	default:
		return "", false
	}
}

func timeoutWait(c chan string, timeout time.Duration) (string, bool) {
	select {
	case m := <-c:
		return m, true
	case <-time.After(timeout):
		return "", false
	}

}

func main() {
	done := make(chan struct{})
	m1 := msgGen("service1", done)
	// m2 := msgGen("service2")
	// m := fanIns(m1, m2)
	for i := 0; i < 5; i++ {
		// m, ok := nonBlockingWait(m1)
		m, ok := timeoutWait(m1, 500*time.Millisecond)
		if ok {
			fmt.Println(m)
		} else {
			fmt.Println("timeout.")
		}
	}
	done <- struct{}{}
	<-done
}
