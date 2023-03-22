package main

import (
	"bufio"
	"fmt"
	"github.com/menah3m/go-tutorial/go-errhandling/fib"
	"os"
)

/*
   @Auth: menah3m
   @Desc:
*/

func tryDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	// err = errors.New("this is a custom err.")
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Path)
			fmt.Println(pathError.Op)
			fmt.Println(pathError.Err)
			fmt.Println(pathError)
		}

		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	writeFile("fib.txt")
	// tryDefer()
}
