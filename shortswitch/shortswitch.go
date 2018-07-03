package main

import (
	"fmt"
	"time"
)

type plugin int

func (p plugin) Example() {
	today := time.Now().Weekday().String()

	switch {
	case today == "Monday":
		fmt.Println("Monday")

	case today == "Tuesday":
		fmt.Println("Tuesday")

	case today == "Wednesday":
		fmt.Println("Wednesday")

	default:
		fmt.Println("The other day")
	}
}

var Plugin plugin
