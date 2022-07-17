package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	data   *Node
	end    *Node
	length int
}

func NewLinkedList(v int) *LinkedList {

	data := Node{value: v}

	return &LinkedList{&data, &data, 1}
}

func (ll *LinkedList) Append(v int) {

	newNode := Node{value: v}

	ll.end.next = &newNode

	ll.end = &newNode

	ll.length += 1

	return

}

func (ll *LinkedList) Prepend(v int) {

	newNode := Node{value: v, next: ll.data}

	ll.data = &newNode

	ll.length += 1

	return
}

func (ll *LinkedList) Insert(index, v int) {

	// don't iterate if insertion is for the first or the end of the list

	if index == 0 {
		ll.Prepend(v)
		return
	} else if index == ll.length {
		ll.Append(v)
		return
	}

	newNode := Node{value: v}

	leader := ll.transverse(index - 1)
	holdPointer := leader.next
	leader.next = &newNode
	newNode.next = holdPointer
	ll.length += 1

	return

}

func (ll *LinkedList) transverse(index int) *Node {

	counter := 0

	currentNode := ll.data

	for counter != index {
		currentNode = currentNode.next
		counter++
	}

	return currentNode
}

func main() {

	newLinkedList := NewLinkedList(10)

	newLinkedList.Append(50)
	newLinkedList.Append(40)
	newLinkedList.Append(70)

	newLinkedList.Prepend(2)

	newLinkedList.Insert(3, 22)

	fmt.Println(newLinkedList.data.next.next.next)

}
