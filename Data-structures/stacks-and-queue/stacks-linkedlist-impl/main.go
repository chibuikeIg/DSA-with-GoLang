package main

import (
	"errors"
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type Stack struct {
	data   *Node
	end    *Node
	length int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) push(value int) {

	newNode := Node{value: value, next: s.data}
	s.data = &newNode

	if s.length == 0 {

		s.end = &newNode

	}

	s.length += 1
}

func (s *Stack) pop() {

	s.data = s.data.next

	s.length -= 1

	if s.length == 0 {
		s.end = nil
	}

}

func (s *Stack) peek() (int, error) {

	if s.length == 0 {
		return 0, errors.New("empty stack")
	}

	return s.data.value, nil
}

func main() {

	newStack := NewStack()

	newStack.push(1)
	newStack.push(2)
	newStack.push(3)

	fmt.Println(newStack.data)

	newStack.pop()
	newStack.pop()
	newStack.pop()

	fmt.Println(newStack.data)

	fmt.Println(newStack.end)

	fmt.Println(newStack.peek())

}
