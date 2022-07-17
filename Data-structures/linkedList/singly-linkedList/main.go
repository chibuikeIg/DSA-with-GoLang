package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	data *Node
	end  *Node
}

func NewLinkedList(v int) *LinkedList {

	data := Node{value: v}

	return &LinkedList{&data, &data}
}

func (ll *LinkedList) Append(v int) {

	newNode := Node{value: v}

	ll.end.next = &newNode

	ll.end = &newNode

}

func (ll *LinkedList) Prepend(v int) {

	newNode := Node{value: v, next: ll.data}

	ll.data = &newNode
}

func main() {

	newLinkedList := NewLinkedList(10)

	newLinkedList.Append(50)
	newLinkedList.Append(40)
	newLinkedList.Append(70)

	newLinkedList.Prepend(2)

	fmt.Println(newLinkedList)

}
