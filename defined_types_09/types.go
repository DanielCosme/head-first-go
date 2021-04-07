package main

import "fmt"

type Liters float64
type Gallons float64

func main() {
	var busFuel Liters
	carFuel := Gallons(10.0)
	busFuel = 240

	fmt.Println(carFuel, busFuel)
	fmt.Printf("%T", carFuel)
}
