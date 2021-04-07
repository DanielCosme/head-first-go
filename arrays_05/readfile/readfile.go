package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var arr [3]float64
	// Open the data file for reading
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Create a new scanner for the file
	scanner := bufio.NewScanner(file)
	var line string
	// Loops until EOF, then scanner.Scan() returns false.
	var i int
	for scanner.Scan() { // Read a line from the file.
		line = scanner.Text()
		arr[i], err = strconv.ParseFloat(line, 64)

		if err != nil {
			log.Fatal(err)
		}

		i++
	}

	err = file.Close() // Close file to free resources
	// If error closing file
	if err != nil {
		log.Fatal(err)
	}
	// If error scanning file
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	var sum, avg float64

	for _, val := range arr {
		sum += val
	}

	avg = sum / float64(len(arr))
	fmt.Println("Sum is", sum)
	fmt.Println("Average is:", avg)
}
