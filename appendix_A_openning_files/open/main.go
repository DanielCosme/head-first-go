package main

import (
	"log"
	"os"
)

func main() {
	// === WRITE ONLY, APPEND, CREATE ===
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE                   // Bitwise
	file, err := os.OpenFile("aardvark.txt", options, os.FileMode(0600)) // octal notation
	check(err)
	_, err = file.Write([]byte("Wow, this is so cool\n"))
	check(err)
	err = file.Close()
	check(err)

	// === READ ONLY ===
	// file, err := os.OpenFile("aardvark.txt", os.O_RDONLY, os.FileMode(0600))
	// check(err)
	// defer file.Close()
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }
	// check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
