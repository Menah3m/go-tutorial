package main

import "fmt"

/*
   @Auth: menah3m
   @Desc:
*/
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	a := adder2(10)

	for i := 0; i < 10; i++ {
		var s int
		s, a = a(i)
		fmt.Printf("sum: %d\n", s)
	}
}
