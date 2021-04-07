package main

import "fmt"

// Defining methods

type Name string

func main() {
	name := Name("Daniel")
	name.sayHi() // name ---> method receiver
}

// (n Name) ---> receiver parameter
func (n Name) sayHi() {
	fmt.Println("Hello,", n)
}
