package main

import "fmt"

//define struct
type Queue struct {
	items []int
}

//enqueue
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

//dequeue
func (q *Queue) Dequeue() int {
	if len(q.items)-1 < 0 {
		fmt.Println("Queue is empty")
		return -1
	}
	r := q.items[0]
	q.items = q.items[1:]
	return r
}

func main() {
	myQueue := Queue{}
	fmt.Println(myQueue)

	myQueue.Enqueue(10)
	myQueue.Enqueue(20)
	myQueue.Enqueue(30)

	fmt.Println(myQueue)

	myQueue.Dequeue()
	myQueue.Dequeue()
	myQueue.Dequeue()
	myQueue.Dequeue()
	fmt.Println(myQueue)
}
