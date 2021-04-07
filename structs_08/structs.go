package main

import (
	"fmt"

	"./magazine"
)

var myStruct struct {
	name     string
	age      int
	money    float64
	isSingle bool
}

// Type definitions allow us to create types of our own,
//		they create a defined type that is based on an underlying type.
// To write a type definition use the type keyword, followed by the name of
//		your new defined type, and then the underlying type to base it on.
type myType struct {
	// Fields go here.
}

func main() {
	myStruct.name = "Daniel"
	myStruct.age = 29
	myStruct.money = 32.89
	myStruct.isSingle = false
	fmt.Println(myStruct)

	var sub *magazine.Subscriber = defaultSubscriber("Daniel")
	applyDiscount(sub)
	printSub(sub)

	sub2 := magazine.Subscriber{Name: "Laura"}
	applyDiscount(&sub2)
	printSub(&sub2)
}

func applyDiscount(s *magazine.Subscriber) {
	s.Rate = 4.99
}

func defaultSubscriber(name string) *magazine.Subscriber {
	var s magazine.Subscriber
	s.Name = name
	s.Rate = 5.99
	s.IsActive = true
	return &s
}

func printSub(sub *magazine.Subscriber) {
	fmt.Println("Name:", sub.Name)
	fmt.Println("Rate:", sub.Rate)
	fmt.Println("Is Active:", sub.IsActive)
}
