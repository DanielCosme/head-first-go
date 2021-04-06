package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	lastName := "Cosme"
	flored := math.Floor(2.75)
	ceiled := math.Ceil(2.75)
	round := math.Round(2.75)
	roundBad := math.Round('a')
	title := strings.Title("head first go")

	fmt.Println("Hello", "Daniel", lastName)
	fmt.Println("Floored", flored)
	fmt.Println("Ceiled", ceiled)
	fmt.Println("Round", round)
	fmt.Println("Round bad", roundBad)
	fmt.Println(title)
}
