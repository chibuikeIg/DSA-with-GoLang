package main

import (
	"fmt"
)

func main() {

	fmt.Println("Answer using interation")

	fmt.Println(findFibonnaciSeqIterative(8))
	fmt.Println("Answer using recursion")

	fmt.Println(findFibonnaciSeqRecursion(8, 0, 1))

	fmt.Println("Answer using iteration by instructor")

	fmt.Println(findFibonnaciSeqIterative1(8))

	fmt.Println("Answer using recursion by instructor")

	fmt.Println(findFibonnaciSeqRecursion1(8))

}

// my solution from exercise

func findFibonnaciSeqIterative(n int) int {

	if n < 2 {

		return n

	}

	prev := 0
	next := 1
	ans := prev + next

	for i := 2; i < n; i++ {

		prev = next
		next = ans
		ans = prev + next

	}

	return ans
}

// my solution from exercise

func findFibonnaciSeqRecursion(num, prev, next int) int {

	ans := prev + next

	if num == 2 {
		return ans
	} else if num < 2 {
		return num
	}

	prev = next

	return findFibonnaciSeqRecursion(num-1, prev, ans)
}

// Instructor's interative solution

func findFibonnaciSeqIterative1(n int) int {

	arr := []int{0, 1}

	for i := 2; i < n+1; i++ {
		arr = append(arr, arr[i-2]+arr[i-1])
	}

	return arr[n]
}

// Instructor's recursive solution

func findFibonnaciSeqRecursion1(num int) int {

	if num < 2 {
		return num
	}

	return findFibonnaciSeqRecursion1(num-1) + findFibonnaciSeqRecursion1(num-2)
}
