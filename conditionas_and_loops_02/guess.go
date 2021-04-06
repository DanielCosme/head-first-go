// challenges players to guess a random number
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var prompt string
	rand.Seed(time.Now().Unix())
	target := rand.Intn(100) + 1
	fmt.Println("I've choosen a random number between 1 and 100")
	fmt.Println("Can you guess it?")
	reader := bufio.NewReader(os.Stdin)

	for guesses := 1; guesses <= 10; guesses++ {
		fmt.Print("Guess my number: ")
		guess, err := reader.ReadString('\n')
		guess = strings.TrimSpace(guess)
		if err != nil {
			log.Fatal(err)
		}
		guessInt, err := strconv.Atoi(guess)
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
