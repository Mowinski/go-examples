package main

import "fmt"

type plugin int

func (p plugin) Example() {
	switch foo := 42; foo {
	case 41:
		fmt.Println("It is 41")

	case 42:
		fmt.Println("It is 42")

	case 43:
		fmt.Println("It is 43")

	default:
		fmt.Println("It is:", foo)
	}
}

var Plugin plugin
