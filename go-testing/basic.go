package main

import (
	"fmt"
	"math"
)

/*
   @Auth: menah3m
   @Desc:
*/

func triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {

	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}
