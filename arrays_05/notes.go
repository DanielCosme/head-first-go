package main

import (
	"fmt"
	"time"
)

func main() {
	var notes [7]string
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"

	for i := 0; i < len(notes); i++ {
		element := fmt.Sprintf("%#v", notes[i])
		fmt.Println("Index:", i, " - Element:", element)
	}

	var dates [3]time.Time
	dates[0] = time.Unix(3245324, 0)
	dates[1] = time.Unix(92387459348057, 0)

	fmt.Println(dates[0])
	fmt.Println(dates[1])
	fmt.Println(dates)

	arr := [3]int{
		1,
		2,
		3,
	}

	fmt.Println(arr)

	for _, value := range notes {
		element := fmt.Sprintf("%#v", value)
		fmt.Println("Element", element)
	}
	fmt.Println()

	pounds := [3]float64{71.8, 56.2, 89.5}
	var sum, avg float64

	for _, val := range pounds {
		sum += val
	}

	avg = sum / float64(len(pounds))
	fmt.Println("Sum is", sum)
	fmt.Println("Average is:", avg)
}
