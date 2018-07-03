package main

import "fmt"

type plugin int

func (p plugin) Example() {
	if foo := 42; foo == 42 {
		fmt.Println("It is true")
	} else {
		fmt.Println("It is", foo)
	}
}

var Plugin plugin
