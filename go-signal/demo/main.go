package main

import (
	"fmt"
	"os"
	"os/signal"
)

/*
   @Auth: menah3m
   @Desc:
*/

func main() {
	// 创建 监听channel 来接收 signal
	// kill 命令 发送的是 SIGTERM
	// Ctrl+C 发送的是 SIGINT
	c := make(chan os.Signal)
	done := make(chan bool)
	signal.Notify(c)
	fmt.Println("start....")
	go func() {
		for true {
			for s := range c {
				fmt.Println("stopped.", s)
				done <- true
			}
		}
	}()

	<-done

}
