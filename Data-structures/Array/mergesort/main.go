package main

import (
	"fmt"
	"sort"
)

func main() {

	arr1 := []int{3, 4, 7, 9}
	arr2 := []int{2, 1, 2, 6}

	fmt.Println("Merge sort 1")
	fmt.Println(mergeSort(arr1, arr2))
	fmt.Println("Merge sort 2")
	fmt.Println(mergeSort2(arr1, arr2))
}

func mergeSort(arr1 []int, arr2 []int) []int {
	for i := 0; i < len(arr2); i++ {
		arr1 = append(arr1, arr2[i])
	}
	sort.Sort(sort.IntSlice(arr1))
	return arr1
}

func mergeSort2(arr1 []int, arr2 []int) []int {
	arr1 = append(arr1, arr2...)
	sort.Sort(sort.IntSlice(arr1))
	return arr1
}
