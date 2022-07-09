package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	fmt.Println("First method:")
	fmt.Println(reverseStr("Hello world peter parker"))
	fmt.Println("Second method:")
	fmt.Println(reverseStr2("Hello world peter parker"))

	fmt.Println("Third method:")
	fmt.Println(reverseStr2("Hello world peter parker"))

}

func reverseStr(str string) string {

	strArr := strings.Split(str, "")
	newStr := ""
	for i := len(strArr) - 1; i >= 0; i-- {
		newStr += strArr[i]
	}
	return newStr
}

func reverseStr2(str string) string {

	newStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		newStr += string(str[i])
	}
	return newStr

}

func reverseStr3(str string) string {

	sliceStr := strings.Split(str, "")
	sort.Sort(sort.Reverse(sort.StringSlice(sliceStr)))

	return strings.Join(sliceStr, "")
}
