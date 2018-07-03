package main

import "fmt"

type plugin int

func (p plugin) Example() {
	foo := 42

	if foo < 42 {
		fmt.Println("It is less than 42")
	} else if foo == 42 {
		fmt.Println("It is 42")
	} else {
		fmt.Println("It is greater")
	}
}

var Plugin plugin
