package main

import "fmt"

func main() {
	numbers := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}

	fmt.Println(bubbleSort(numbers)) // 0 1 2 4 5 6 44 63 87 99 283
}

func bubbleSort(nums []int) []int {

	len := len(nums)

	for i := 0; i < len; i++ {

		for j := 0; j < len; j++ {

			index := j + 1

			if index < len && nums[j] > nums[index] {

				temp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp

			}

		}
	}

	return nums
}
