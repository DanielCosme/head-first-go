package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var args = os.Args[1:]
	var sum, avg float64

	for _, arg := range args {
		number, err := strconv.ParseFloat(arg, 64)

		if err != nil {
			fmt.Println(err)
			continue
		}

		sum += number
	}

	avg = sum / float64(len(args)) // Potential division by 0.
	fmt.Printf("Average is %0.1f\n", avg)
	fmt.Println("Sum is:", sum)
	fmt.Printf("%v is of type %T", math.NaN(), math.NaN())
}
