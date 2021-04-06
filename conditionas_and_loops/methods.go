package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func main() {
	// The time package has a Time type that represents a date and time.
	// Each time.Time type value has a Year method that returns that year.
	// year, month, day - hour, minute, second
	var now time.Time = time.Now()
	year := now.Year()

	fmt.Println("Full date", now)
	fmt.Println("Year", year)
	fmt.Println("Type of year", reflect.TypeOf(year))
	fmt.Println("Month", now.Month())
	fmt.Println("Day", now.Day())
	fmt.Println()

	// The strings package has a Replacer that can search throught a string for
	// a substring and replace each ocurrence of that substring with another
	// substring .
	var broken string = "G# R#cks"
	//replacer := strings.NewReplacer("#", "o")
	fixed := strings.NewReplacer("#", "o").Replace(broken)
	fmt.Println("Broken", broken)
	fmt.Println("Fixed", fixed)

}
