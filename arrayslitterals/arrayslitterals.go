package main

import "fmt"

type plugin int

func (p plugin) Example() {
	numbers := [5]int{1, 2, 3, 4, 5}
	otherNumber := [...]string{"Hello", "World"}

	fmt.Println(numbers)
	fmt.Println(otherNumber)
}

var Plugin plugin
