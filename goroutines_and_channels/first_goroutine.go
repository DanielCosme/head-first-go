package main

import (
	"fmt"
	"time"
)

func main() {
	go run("a", 50)
	go run("b", 50)
	time.Sleep(time.Second)
	fmt.Println("End Main")
}

func run(l string, t int) {
	for i := 0; i < t; i++ {
		fmt.Print(l)
	}
	fmt.Println()
}
