package queue

/*
   @Auth: menah3m
   @Desc:
*/

// An FIFO queue.
type Queue []int

// Push pushes the element into the queue
//       e.g q.Push(123)
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

// Pop pops element from head.
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// IsEmpty Returns false if queue is empty
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
