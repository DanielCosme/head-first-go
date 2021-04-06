package main

import (
	"fmt"
	"reflect"
)

func main() {
	// String literals
	fmt.Println("Hello, \nwith new line character")
	fmt.Println("Hello, \t with tab charcter")
	fmt.Println("Hello, \"\" with quotes")
	fmt.Println("Hello, \\ with a backslash", "\n")

	// Runes
	// Used to represent single characters
	// Rune literals are sorrounded by single quotation marks ''
	// Runes are kept as their nummeric representation from the
	// unicode stantard
	fmt.Println("Rune of A", 'A')
	fmt.Println("Rune of a", 'a')
	fmt.Println("Rune of newline char", '\n')
	fmt.Println("Rune of tab char", '\t')
	fmt.Println("Rune of double quote", '"')
	fmt.Println("Rune of backslash", '\\')
	fmt.Println()

	// Booleans
	fmt.Println("Boolean literal", true)
	fmt.Println("Boolean literal", false)
	fmt.Println()

	// Numbers
	fmt.Println("Number literal", 12)
	fmt.Println("Floating Point literal", 12.1)
	fmt.Println()

	// The type of any value can be seen using the TypeOf function
	// from the reflect package
	fmt.Println("12 is of type:", reflect.TypeOf(12))
	fmt.Println("12.0 is of type:", reflect.TypeOf(12.0))
	fmt.Println("\"Daniel\" is of type:", reflect.TypeOf("Daniel"))
	fmt.Println("\"true\" is of type:", reflect.TypeOf(true))
	fmt.Println("")

	// Declaring variables
	var quantity int
	var length, width float64 = 1.0, 2.4
	var str, f = "str", 1.1
	var name string
	fmt.Println("Quantity is:", quantity)
	fmt.Println("Length is:", length, "Width is:", width)
	fmt.Println("Name is:", name)
	fmt.Println("str", str, "f", f)
	fmt.Println()

	quantity = 14
	length = 12.3
	width = 13.4
	name = "Daniel"
	fmt.Println("Quantity is:", quantity)
	fmt.Println("Length is:", length, "Width is:", width)
	fmt.Println("Name is:", name)

	// Short variable declarations
	q := 1
	w, l := 1.0, 3.3
	n := "Dan"

	fmt.Println("q", q, "l", l, "w", w, "n", n)
	fmt.Println()
	// Exported functions/variables are Capitalized (camel case)
	// Non-exported functions/variables are pascalCase.

	// Type conversions
	// When doing math and comparison operation both operands must be of
	// the same type
	var myInt int = 34
	var myFloat float64 = 2.0
	res := myInt * int(myFloat)
	fmt.Println(res)
	fmt.Println()

}
