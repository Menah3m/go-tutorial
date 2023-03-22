package queue

import "fmt"

/*
   @Auth: menah3m
   @Desc:
*/
func ExampleQueue_Push() {
	q := Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	// Output:
	// 1
	// 2
	// false
	// 3
	// true
}
func ExampleQueue_Pop() {

}
func ExampleQueue_IsEmpty() {

}
