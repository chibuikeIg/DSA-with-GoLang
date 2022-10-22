package main

import "fmt"

func main() {

	numbers := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}

	fmt.Println(selectionSort(numbers)) // 0 1 2 4 5 6 44 63 87 99 283
}

func selectionSort(numbers []int) []int {

	len := len(numbers)

	for i := 0; i < len; i++ {

		lowestNum := numbers[i]
		index := i

		for j := i + 1; j < len; j++ {

			if lowestNum > numbers[j] {
				lowestNum = numbers[j]
				index = j
			}

		}

		numbers[index] = numbers[i]

		numbers[i] = lowestNum

	}

	return numbers
}
