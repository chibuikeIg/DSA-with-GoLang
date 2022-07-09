package main

import "fmt"

func main() {

	numbers := []int{1, 5, 6, 7, 5, 3}

	fmt.Println(findSumMatchingPair2(numbers, 2))

	// possible improvements and error checks to be made
	// 1. Check if collection is empty
	// 2. Check the length of collection is greater than 1
	// 3. Maybe check for negative values
}

// interview question:
// Given a collection of numbers, take the collection of numbers and find a matching pair that is equal to any given sum

// Possible questions I can ask the interviewer
/// 1. What should the function return when it finds the matching pair? true or false or the matched pair?
//  2. Can numbers repeat itself in the collection?
//  3. Can we assume the length of the collection will always be greater than 1?
//  4. What should I put into consideration? Time complexity or Space complexity or just the result?

// for example
// given
// first input/param: numbersArr = [1, 3, 5, 6, 7]
// second input/param: sum = 6
// output: true or false
// the matching pair in the numbers collection is 1 + 5 = 6

//naive solution

func findSumMatchingPair(numCollection []int, sum int) bool {

	for i := 0; i < len(numCollection); i++ {
		for j := 0; j < len(numCollection); j++ {
			if foundPairSum := numCollection[i] + numCollection[j]; foundPairSum == sum {
				return true
			}
		}
	}

	return false
}

// Big O FOR time complexity = O(n^2) //  horrible
// Big O for space complexity = O(1) // GOOD

// second solution

func findSumMatchingPair2(numCollection []int, sum int) bool {

	// create or initialize an empty map (or associative array) to store items in the collections whose sum difference wasn't found

	newNumCollection := make(map[int]bool)

	// loop through the given collection and check if the difference is available in the new collection and if yes return true

	for i := 0; i < len(numCollection); i++ {
		diff := sum - numCollection[i]
		if _, ok := newNumCollection[diff]; ok {
			return true
		}
		newNumCollection[numCollection[i]] = true
	}

	return false
}

// Big O FOR time complexity = O(n) //  Good
// Big O for space complexity = O(n^2)
