package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Concurrency patterns
// Multiplexing
// fanIn function lets talk to wohever is ready to talk

func main() {
	quit := make(chan string)
	jane := boring("Jane!")
	daniel := boring("Daniel!")
	other := boring("Timeout")

	c := fanInSelect(jane, daniel)
	timeOut(other, quit)
	for i := 0; i < 5; i++ {
		fmt.Println(<-c, "\n")
	}

	quit <- "Bye!"
	fmt.Printf("Timeout says: %q\n", <-quit)
}

// Timeout using select
func timeOut(c <-chan string, quit chan string) {
	timeout := time.After(8 * time.Second)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("You are to slow")

		case <-quit:
			fmt.Println("CLEANING")
			quit <- "See you!"
		}
	}

}

func fanInSelect(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

// Channels are first-class values, just like strings or integers.
func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
