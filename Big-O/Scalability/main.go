package main

import (
	"fmt"
	"time"
)

func main() {

	// nemo := []string{"nemo"}

	// everyone := []string{"Gregory", "Julios", "David", "Nemo", "Limo", "Zimo", "Dimo", "Gimo", "Nino", "Teno"}
	largeArr := GenerateArr(1000, "Nemo")
	FindNemo(largeArr)
}

func FindNemo(array []string) {

	t0 := time.Now()

	for i := 0; i < len(array); i++ {

		if array[i] == "Nemo" {

			fmt.Println("Found nemo")

		}

	}

	t1 := time.Now()
	diff := t1.Sub(t0).Milliseconds()

	fmt.Println("Call to find nemo took ", diff, " Milliseconds")
}

func GenerateArr(size int, value string) []string {

	var largeArr []string

	for i := 0; i < size; i++ {
		largeArr = append(largeArr, value)
	}

	return largeArr
}
