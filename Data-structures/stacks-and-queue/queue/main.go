package main

import "fmt"

type Node struct {
	value any
	next  *Node
}

type Queue struct {
	first  *Node
	last   *Node
	length int
}

func NewQueue() *Queue {

	return &Queue{}

}

func (q *Queue) enqueue(value any) {

	newNode := Node{value: value}
	nodePointer := &newNode

	if q.length == 0 {

		q.first = nodePointer
		q.last = nodePointer

	} else {

		q.last.next = nodePointer
		q.last = nodePointer

	}

	q.length += 1

}

func (q *Queue) dequeue() {

	q.first = q.first.next
	q.length -= 1

	if q.length == 0 {
		q.last = nil
	}

}

func (q *Queue) peek() *Node {

	return q.first

}

func main() {

	newQueue := NewQueue()

	newQueue.enqueue("Joy")
	newQueue.enqueue("Matt")
	newQueue.enqueue("Pavel")
	newQueue.enqueue("Samir")

	fmt.Println(newQueue)

	newQueue.dequeue()

	fmt.Println(newQueue.peek())

}
