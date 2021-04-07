package main

import (
	"fmt"
)

// An error value in golang is any value that has a method Error that returns
//		a string.

func paintNeeded(width float64, heigth float64) (result float64, err error) {
	// The ability to create a value that represents an error
	// The hability to return an additional value that represents the error
	if width < 0 {
		err = fmt.Errorf("Width cannot be negative: %.2f", width)
	}
	if heigth < 0 {
		err = fmt.Errorf("heigth cannot be negative: %.2f", heigth)
	}

	if err == nil {
		area := width * heigth
		result = area / 10
	}

	return result, err
}

func main() {
	var total, amount float64
	// amount = paintNeeded(4.2, 3.0)
	// total += amount
	// fmt.Printf("%.2f Liters ares needed.\n", amount)

	amount, err := paintNeeded(4.2, 3.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%.2f Liters ares needed.\n", amount)
		total += amount
	}

	amount, err = paintNeeded(-5.2, -3.5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%.2f Liters ares needed.\n", amount)
	}
	total += amount

	fmt.Printf("%.2f Total needed.\n", total)

}
