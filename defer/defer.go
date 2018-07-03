package main

import "fmt"

type plugin int

func (p plugin) Example() {
	defer fmt.Println("last word")
	defer fmt.Println("last but one")

	fmt.Println("first")
}

var Plugin plugin
