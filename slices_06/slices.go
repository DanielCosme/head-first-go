package main

import "fmt"

func main() {
	// A slice is a collection that has variable lengths, we can add new values
	//  	as we go.
	// To create a slice you leave the size of the array [] empty.
	var mySlice []string // slice
	var myArr [2]string  // array

	// Declaring a slice variable does not make it a slice.
	// For that we need to invoke the make function.
	notes := make([]int, 7)

	fmt.Println(len(mySlice))
	fmt.Println(len(myArr))
	fmt.Println(len(notes))
	fmt.Println(notes)
	fmt.Println()

	// Slice literals
	s := []int{1, 2, 3}
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", myArr)
	fmt.Println()

	// Slices are built on top of arrays
	// One cannot understand how slices work without understanding arrays
	// Every slice is build on top of an underlying array. The array is what
	//		actually holds the data. The slice is a mere view into some elements
	//		of the array (or all).
	// The slice operator

	underlyingArray := [5]string{"a", "b", "c", "d", "e"}
	slice1 := underlyingArray[1:4]
	slice2 := underlyingArray[:3] // Initial index defaults to 0
	slice3 := underlyingArray[1:] // Last index defaults to the end of the array
	slice4 := slice3
	fmt.Println("Underlying Array", underlyingArray)
	fmt.Println("Slice 1", slice1)
	fmt.Println("Slice 2", slice2)
	fmt.Println("Slice 3", slice3)
	slice4[1] = "Z"
	slice4[2] = "Y"
	fmt.Println("Slice 4", slice4)
	fmt.Println("Underlying Array", underlyingArray)
	fmt.Println("Slice 3", slice3)
	fmt.Println("Slice 1", slice1)
	fmt.Println()

	// Append to slice.
	slice5 := []int{1, 2, 3, 4}
	fmt.Println("Slice 5", slice5, len(slice5))
	slice := append(slice5, 5)
	fmt.Println("Slice", slice, len(slice))
	slice = append(slice, 6, 7)
	fmt.Println("Slice", slice, len(slice))

}
