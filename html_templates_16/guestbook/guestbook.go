package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Respond to requests for the main guestbook page.
// Format the response using HTML.
// Fill the HMTL page with signatures. (Read data from file)

// Set up a form for adding a new signature.
// Save submitted signatures into file.

type GuestBook struct {
	SignatureCount int
	Signatures     []string
}

func main() {
	http.HandleFunc("/guestbook", guestbookHandler)
	http.HandleFunc("/signature", signatureHander)
	http.HandleFunc("/signature/new", newSignatureHandler)

	err := http.ListenAndServe(":3000", nil)
	log.Fatal(err)
}

func newSignatureHandler(rw http.ResponseWriter, r *http.Request) {
	signature := r.FormValue("signature")

	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("signatures.txt", options, os.FileMode(0600))
	check(err)

	_, err = fmt.Fprintln(file, signature)
	check(err)

	err = file.Close()
	check(err)

	http.Redirect(rw, r, "/guestbook", http.StatusFound)
}

func guestbookHandler(rw http.ResponseWriter, r *http.Request) {
	guestbookTemplate, err := template.ParseFiles("guestbook.html")
	check(err)
	signaturesFromFile := getStrings("signatures.txt")
	guestBook := GuestBook{
		SignatureCount: len(signaturesFromFile),
		Signatures:     signaturesFromFile,
	}

	err = guestbookTemplate.Execute(rw, guestBook)
	fmt.Println(err)
}

func signatureHander(rw http.ResponseWriter, r *http.Request) {
	htmlTemplate, err := template.ParseFiles("signature.html")
	check(err)
	err = htmlTemplate.Execute(rw, nil)
	fmt.Println(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getStrings(fileName string) (result []string) {
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	check(scanner.Err())

	return
}
