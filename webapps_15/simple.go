package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = "localhost:3000"

func main() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/hi", viewHandler)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func defaultHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Default")

	_, err := writer.Write(message)
	fmt.Println("Number of bytes", request)
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello, Web! JJ")

	l, err := writer.Write(message)
	fmt.Println("Number of bytes", l)
	if err != nil {
		log.Fatal(err)
	}
}
