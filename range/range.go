package main

import "fmt"

type plugin int

func (p plugin) Example() {
	numbers := []int{1, 2, 3, 4, 5}
	for i, v := range numbers {
		fmt.Print(i, ": ")
		fmt.Println(v * v)
	}
}

var Plugin plugin
