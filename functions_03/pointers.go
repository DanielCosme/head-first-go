package main

import (
	"fmt"
	"reflect"
)

// Golang is pass by value.

func main() {
	// One can get the address of a variable using ampersand (&)
	// Ampersand is the "address of" operator in golang.
	value := 6
	fmt.Println("Value", value)
	fmt.Println("Memory address", &value)

	// Pointer types
	// The type of a pointer is written with a * followed by the type of the
	// variable that it points to.
	// The thype of a pointer to an int is *int --> pointer to int

	var myInt int = 10
	var myFloat float64 = 34.834
	t := reflect.TypeOf(&myInt)
	fmt.Println("type of", t)
	fmt.Printf("Type of myInt is %T\n", &myInt)
	fmt.Printf("Type of myFloat is %T\n", &myFloat)

	// We van declare variables that hold pointers.
	// A pointer variable can only hold pointers to one type of value (type)
	var myIntPointer *int // Variable that holds a pointer to an int.
	myIntPointer = &myInt // Assign a pointer to the variable.
	fmt.Println("My int pointer (reference) value is", myIntPointer)
	fmt.Println("My int pointer value is", *myIntPointer)
	fmt.Println()

	myFloatPointer := &myFloat // Assign a pointer to the variable (short-hand)
	fmt.Println("My float pointer (reference) value is", myFloatPointer)
	fmt.Println("My float pointer value is", *myFloatPointer)
	fmt.Println()

	// Getting or changing the value of a pointer
	// operator (*) reads --> get value at pointer (memory address).
	// This operator can also be used in order to update the value of a pointer

	fmt.Println("Current value of myInt:", myInt)
	*myIntPointer = 8
	fmt.Println("Value myIntPointer after de-referencing", *myIntPointer)
	fmt.Println("Value of myInt after using the * operator", myInt)
	fmt.Println()

	var myFloatOtherPointer *float64 = createPointer()
	fmt.Println("Value of pointer", myFloatOtherPointer)
	fmt.Println("Derefenced value", *myFloatOtherPointer)
	fmt.Println()

	var myBool bool
	var myBoolPointer *bool = &myBool // Creating a reference.
	printPointer(myBoolPointer)
	fmt.Println()

	fmt.Println("Current value of myInt:", myInt)
	double(&myInt)
	fmt.Println("Value of myInt after invoking double func", myInt)
}

// One can also pass pointers to functions as arguments.
// Just specify that the type of the parameters is that of pointer.
func printPointer(myBoolPointer *bool) {
	fmt.Println(myBoolPointer) // Print the value at the pointer that gets passed in
}

func createPointer() *float64 {
	// It is posible to return pointers from functions, one just need to declare
	//		that the function's return is of type pointer
	var f float64
	return &f
}

func double(number *int) {
	n := number
	*n *= 2
}
