package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) DeleteWithvalue(value int) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	prev := l.head
	for prev.next.data != value {
		if prev.next.next == nil {
			return
		}
		prev = prev.next
	}

	prev.next = prev.next.next
	l.length--
}

//Insert
func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

//Output list
func (l linkedList) PrintList() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)

		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func main() {
	myList := linkedList{}

	node1 := &node{data: 10}
	node2 := &node{data: 20}
	node3 := &node{data: 40}
	node4 := &node{data: 50}
	node5 := &node{data: 60}
	node6 := &node{data: 70}

	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.prepend(node6)

	myList.PrintList()
	myList.DeleteWithvalue(40)
	myList.PrintList()

	myList.DeleteWithvalue(100)
	myList.DeleteWithvalue(10)
	emptyList := linkedList{}
	emptyList.DeleteWithvalue(333)

}
