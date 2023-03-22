package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

/*
   @Auth: menah3m
   @Desc:  go 执行 linux 命令
*/

func main() {
	// 创建cmd实例
	cmd := exec.Command("ls", "-a")

	// 创建一个标准输出管道
	output, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 执行命令
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return
	}
	// 读取结果
	bytes, err := ioutil.ReadAll(output)
	fmt.Printf("%s", bytes)
}
