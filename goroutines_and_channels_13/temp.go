package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Channels is how we communicate between goroutines.
// They allow to send messages from one goroutine to another.
// They ensure the sending goroutine has sent the value before the receving
//		goroutine attempts to use it.
// To demostrate channels we need to:
// Create a channel.
// Write a function that receives a channel a parameter.
//		This function runs in a separate goroutine, and its used to send values
//		over the channel.
// Retreive sent values in original goroutine.

func main() {
	// Each channel only carries values of a particular type.
	myChannel := make(chan Url)

	// To send a value to a channel, you use the <- operator
	//		channel <- value
	// We also use the <- operator to recive values from a channel.
	// However, the positioning is different, you place the arrow to the left to
	//		the channel you are receiving from.
	// <- channel

	urls := []string{
		"https://example.com",
		"https://google.com",
		"https://instagram.com",
		"https://facebook.com",
		"https://youtube.com",
		"https://golang.org",
		"https://apple.com",
		"https://danicos.me",
	}
	for _, url := range urls {
		go responseSize(url, myChannel) // Concurrent operation
	}
	for i := 0; i < len(urls); i++ {
		fmt.Print(<-myChannel)
	}
}

type Url struct {
	url string
	len int
}

func (u Url) String() (res string) {
	res += u.url + " - "
	res += fmt.Sprint(float64(u.len)/1000.0, " MB", "\n")
	//res += fmt.Sprint(u.len*8, " Bits", "\n\n")
	return
}

func responseSize(url string, channel chan Url) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	resultUrl := Url{url, len(body)}
	channel <- resultUrl
}
