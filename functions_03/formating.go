package main

import "fmt"

func main() {
	// Formating output with Printf and Sprintf
	fmt.Println("About one-third", 1.0/3.0)
	// Printf = print with formating, takes a string and inserts one or more
	//   values into it.
	fmt.Printf("About one third: %0.2f\n", 1.0/3.0)
	// Sprintf, returns a formated string instead of printing it.
	formated := fmt.Sprintf("About one third: %0.2f\n", 1.0/3.0)
	fmt.Print(formated)
	fmt.Println()

	// Value widths
	// if value is less than the width specified it will padd it with spaces

	// The first argument to Printf is a string that will be used to format the
	//		output.
	// Any percentage (%) signs will be treated as the starting of the formating
	//		verb
	// Formating verbs --> a section of the string that will be substituted with
	//		a value in a particular format.
	// The remaining arguments are used as values for those verbs.
	// The letter following the percent sign indicates which verb to use.

	fmt.Printf("A float %0.1f\n", 3.84) // 3.8
	fmt.Printf("An integer %d\n", 34)
	s := "Hello World!"
	fmt.Printf("A string, %s\n", s)
	fmt.Printf("A boolean %t\n", true)
	fmt.Printf("A boolean %t\n", false)
	fmt.Printf("Several values %v %v %v %v\n", 1.2, "\t", "true", '3')
	fmt.Printf("Several values %#v %#v %#v %#v\n", 1.2, "\t", "true", '3')
	fmt.Printf("Types %T %T %T\n", s, 2.3, true)
	fmt.Printf("A percent sign: %%\n")

	fmt.Printf("\n%12s | %s\n", "Product", "Cost in Cents")
}
