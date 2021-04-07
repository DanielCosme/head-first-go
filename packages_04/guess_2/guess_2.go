// challenges players to guess a random number
package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/danielcosme/keyboard"
)

func main() {
	var prompt string
	rand.Seed(time.Now().Unix())
	target := float64(rand.Intn(100) + 1)
	fmt.Println("I've choosen a random number between 1 and 100")
	fmt.Println("Can you guess it?")

	for guesses := 1; guesses <= 10; guesses++ {
		fmt.Print("Guess my number: ")
		guessInt, err := keyboard.GetFloat()
		if err != nil {
			log.Fatal(err)
		}

		if guessInt < target {
			prompt = "Ooops. Your guess was LOW"
		} else if guessInt > target {
			prompt = "Ooops. Your guess was HIGH"
		} else {
			prompt = "Good job! You gessed it!"
			fmt.Println(prompt)
			break
		}

		if guesses == 10 {
			fmt.Println("Sorry. You didin't guessed my number. It was:", target)
			break
		}

		fmt.Println(prompt)
		fmt.Println("Number of guesses left", 10-guesses)
		fmt.Println()
	}
}
