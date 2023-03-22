package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*
   @Auth: menah3m
   @Desc:
*/

func main() {

	// 创建监听信号的channel（订阅）
	c := make(chan os.Signal)
	flag := true

	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		// 当接收到对应信号时，执行处理函数中的内容
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				t := time.Now()
				fmt.Println("program stopping")
				flag = false
				GracefulQuit()
				fmt.Println("program stopped.. Signal type:", s)
				fmt.Println(time.Now().Sub(t))
				os.Exit(0)
			case syscall.SIGUSR1:
				fmt.Println("usr1 signal:", s)
			case syscall.SIGUSR2:
				fmt.Println("usr2 signal:", s)
			default:
				fmt.Println("other signal:", s)

			}
		}
	}()

	fmt.Println("program start...")
	sum := 0
	// 处理函数在执行时，主函数停止运行，可以设置个flag
	for flag {
		sum++
		fmt.Println("Sum:", sum)
		time.Sleep(time.Second)
	}
	wg.Wait()
}

// 注册处理函数
func GracefulQuit() {
	fmt.Println("start quit...")
	fmt.Println("cleaning...")
	time.Sleep(3 * time.Second)
	fmt.Println("quit.")

}
