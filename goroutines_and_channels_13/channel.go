package main

import (
	"fmt"
	"time"
)

// The main function is the receiver goroutine.
func main() {
	var myChannel chan string = make(chan string)
	go send(myChannel)
	reportNap("Receiving Goroutine", 5)
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
}

// The send function is the sender goroutine.
func send(c chan string) {
	reportNap("Sender Goroutine", 2)
	fmt.Println("*** Sending Value ***")
	c <- "A"
	fmt.Println("*** Sending Value ***")
	c <- "B"
}

func reportNap(user string, delay int) {
	fmt.Println(user, "Goes to Sleep")
	for count := 0; count < delay; count++ {
		fmt.Println(user, "Sleeping")
		time.Sleep(time.Second)
	}
	fmt.Println(user, "Wakes up")
}
