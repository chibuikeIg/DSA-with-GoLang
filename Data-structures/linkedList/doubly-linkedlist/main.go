package main

import "fmt"

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type DoublyLinkedList struct {
	data   *Node
	end    *Node
	length int
}

func NewDoublyLinkedList(v int) *DoublyLinkedList {

	data := Node{value: v}

	return &DoublyLinkedList{&data, &data, 1}
}

func (ll *DoublyLinkedList) Append(v int) {

	newNode := Node{value: v, prev: ll.end}

	ll.end.next = &newNode

	ll.end = &newNode

	ll.length += 1

	return

}

func (ll *DoublyLinkedList) Prepend(v int) {

	newNode := Node{value: v, next: ll.data}

	ll.data.prev = &newNode

	ll.data = &newNode

	ll.length += 1

	return
}

func (ll *DoublyLinkedList) Insert(index, v int) {

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
	follower := leader.next
	leader.next = &newNode
	newNode.prev = leader
	newNode.next = follower
	follower.prev = &newNode
	ll.length += 1

	return

}

func (ll *DoublyLinkedList) transverse(index int) *Node {

	counter := 0

	currentNode := ll.data

	for counter != index {
		currentNode = currentNode.next
		counter++
	}

	return currentNode
}

func (ll *DoublyLinkedList) Remove(index int) {

	if index == 0 {

		ll.data = ll.data.next
		ll.data.prev = nil
		ll.length -= 1

		return
	}

	headNode := ll.transverse(index - 1)

	follower := headNode.next.next

	headNode.next = follower

	if index != ll.length-1 {

		follower.prev = headNode
	}

	ll.length -= 1

}

func (ll *DoublyLinkedList) PrintList() {

	list := make([]int, ll.length)

	counter := 0

	currentNode := ll.data

	for counter != ll.length {

		list[counter] = currentNode.value

		currentNode = currentNode.next

		counter++
	}

	fmt.Println(list)

}

func main() {

	newLinkedList := NewDoublyLinkedList(10)

	newLinkedList.Append(50)
	newLinkedList.Append(40)
	newLinkedList.Append(70)

	newLinkedList.Prepend(2)

	newLinkedList.Insert(3, 22)

	newLinkedList.Insert(2, 12)

	newLinkedList.Remove(5)

	newLinkedList.PrintList()

}
