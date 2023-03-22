package main

import (
	"fmt"
)

/*
   @Auth: menah3m
   @Desc:
*/
var LastOccurred = make([]int, 0xffff)

func main() {
	// a, b := 3, 4
	// fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	// fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
	// swap(&a, &b)
	// fmt.Println(a, b)
	// var arr1 [5]int
	// arr2 := [3]int{1, 2, 3}
	// arr3 := [...]int{2, 3, 4, 56, 3}
	// grid := [4][3]int{}
	// fmt.Println(arr1, arr2, arr3)
	// fmt.Println(grid)
	// for _, v := range arr3 {
	// 	fmt.Println(v)
	// }

	// s := make([]int, 16)
	// for i := 0; i < 100; i++ {
	// 	printSlice(s)
	//
	// 	s = append(s, 2*i+1)
	// }

	// fmt.Println(GetMaxLengthOfNonRepeatingSubStr("b"))
	// fmt.Println(GetMaxLengthOfNonRepeatingSubStr(""))
	fmt.Println(GetMaxLengthOfNonRepeatingSubStr("abcdefghi"))
	// fmt.Println(GetMaxLengthOfNonRepeatingSubStr("abcabcbb"))
	// fmt.Println(GetMaxLengthOfNonRepeatingSubStr("bbbbb"))
	fmt.Println(GetMaxLengthOfNonRepeatingSubStrByRune("一二三yy"))

	// s := "yes你好世界"
	// fmt.Println(len(s))
	// for i, v := range s {
	// 	fmt.Printf("%d %X\n", i, v)
	// }
	// for _, b := range []byte(s) {
	// 	fmt.Printf("%X\n", b)
	// }
	// fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// bytes := []byte(s)
	// for len(bytes) > 0 {
	// 	ch, size := utf8.DecodeRune(bytes)
	// 	bytes = bytes[size:]
	// 	fmt.Printf("%c ", ch)
	// }
	// fmt.Println()
	//
	// for i, ch := range []rune(s) {
	// 	fmt.Printf("%d %c\n", i, ch)
	// }

}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func printSlice(s []int) {
	fmt.Println("len=", len(s), "cap=", cap(s))
}

func GetMaxLengthOfNonRepeatingSubStr(s string) int {
	start := 0
	maxLength := 0
	LastOccurred := make(map[byte]int)
	for i, v := range []byte(s) {
		lastI, ok := LastOccurred[v]
		if ok && lastI >= start {
			start = LastOccurred[v] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		LastOccurred[v] = i
	}
	return maxLength
}

// 国际版 优化前
// func GetMaxLengthOfNonRepeatingSubStrByRune(s string) int {
// 	start := 0
// 	maxLength := 0
// 	LastOccurred := make(map[rune]int)

// 	for i, v := range []rune(s) {
// 		lastI, ok := LastOccurred[v]
// 		if ok && lastI >= start {
// 			start = LastOccurred[v] + 1
// 		}
// 		if i-start+1 > maxLength {
// 			maxLength = i - start + 1
// 		}
// 		LastOccurred[v] = i
// 	}
// 	return maxLength
// }

// 优化后
func GetMaxLengthOfNonRepeatingSubStrByRune(s string) int {
	start := 0
	maxLength := 0

	for i := range LastOccurred {
		LastOccurred[i] = -1
	}

	for i, v := range []rune(s) {

		if lastI := LastOccurred[v]; lastI != -1 && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		LastOccurred[v] = i
	}
	return maxLength
}
