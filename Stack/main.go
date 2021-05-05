package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

func (s *Stack) Pop() int {
	l := len(s.items) - 1

	if l < 0 {
		fmt.Println("stack is empty")
		return -1
	}

	r := s.items[l]
	s.items = s.items[:l]
	return r
}

func main() {
	myStack := Stack{}

	fmt.Println(myStack)

	myStack.Push(10)
	myStack.Push(20)
	myStack.Push(30)
	myStack.Push(40)
	fmt.Println(myStack)

	myStack.Pop()
	myStack.Pop()
	myStack.Pop()
	myStack.Pop()
	myStack.Pop()

}
