package main

import "fmt"

type plugin int

func (p plugin) Example() {
	fmt.Println("Hello, World")
}

var Plugin plugin
