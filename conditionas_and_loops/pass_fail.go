// Reports whether a grade is passing of failing.
package main // Main package because the package is an executable.

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var status string
	// Getting a grade from the user.
	// Prompt user for a grade, save grade in a variable.
	fmt.Print("Enter a grade: ")
	// Buffered reader that gets text from keyboard.
	reader := bufio.NewReader(os.Stdin)
	// Return everything the user has typed, until delimiter reached.
	// In go functions/methods can return any number of values.
	// Blank identifier in order to ignore (discard) the error.
	input, err := reader.ReadString('\n')
	if err != nil {
		// Log error and exit program.
		log.Fatal(err)
	}

	// Converting strings to numbers.
	// Need to strip the '\n' character at the end of the string
	input = strings.TrimSpace(input)

	// Convert the string into a number
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}

	if grade == 100 {
		status = "Perfect!"
	} else if grade >= 60 {
		status = "Passing"
	} else {
		status = "Failing"
	}

	fmt.Println("A grade of", grade, "is", status)
}
