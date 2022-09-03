package main

import "fmt"

type Stack struct {
	array []string
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) push(value string) {

	s.array = append(s.array, value)

}

func (s *Stack) pop() {

	s.array = s.array[:len(s.array)-1]

}

func (s *Stack) peek() string {

	if len(s.array) == 0 {

		return ""
	}

	return s.array[len(s.array)-1]

}

func main() {

	newStack := NewStack()

	newStack.push("Google")
	newStack.push("Discord")
	newStack.push("Udemy")

	fmt.Println(newStack)

	newStack.pop()
	newStack.pop()
	newStack.pop()

	fmt.Println(newStack.peek())

	fmt.Println(newStack)

}
