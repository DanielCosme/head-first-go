package main

import (
	"fmt"
)

// A variadic function is one that can be called with varying number of
//		arguments
// To make a function variadic we use an elipsis before the type of the last
//		argument.

func main() {
	slice := []int{343, 434, 434, 546}
	arr := [4]int{44, 33, 22, 11}
	toPrint(1, 2, 3, 4, 5, 6)
	toPrint(99, 1, 0)
	toPrint()
	// Passing slices to variadic functions.
	toPrint(slice...)
	fmt.Println(slice)

	s := arr[0:]
	fmt.Println(arr)
	mutate(s...)
	fmt.Println(arr)
}

func mutate(arr ...int) {
	arr[0] = 0
}

// Variadic function.
func toPrint(numbers ...int) {
	if len(numbers) == 0 {
		return
	}

	for _, v := range numbers {
		fmt.Println(v)
	}

	fmt.Println()
}
