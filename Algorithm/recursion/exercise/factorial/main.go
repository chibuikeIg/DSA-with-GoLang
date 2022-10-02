package main

import "fmt"

func main() {

	fmt.Println("Answer from recursion")
	fmt.Println(findFactorialRecursive(5))
	fmt.Println("Answer from iteration")
	fmt.Println(findFactorialIterative(5))
}

func findFactorialRecursive(n int) int {

	answer := n

	if n == 1 {
		return answer
	}

	return answer * findFactorialRecursive(n-1)

}

func findFactorialIterative(n int) int {

	answer := n * (n - 1)

	for i := n - 2; i >= 1; i-- {
		answer *= i
	}

	return answer
}
