package main

import (
	"errors"
	"fmt"
	"log"
)

type Date struct {
	year  int
	month int
	day   int
}

func main() {
	date := Date{}
	err := date.setYear(2019)
	if err != nil {
		log.Fatal(err)
	}

	date.setDay(27)
	date.setMonth(5)
	fmt.Println(date)
}

func (d *Date) setYear(y int) error {
	if y < 1 {
		return errors.New("Invalid year")
	}
	d.year = y
	return nil
}

func (d *Date) setMonth(m int) {
	d.month = m
}

func (d *Date) setDay(day int) {
	d.day = day
}
