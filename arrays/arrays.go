package main

import "fmt"

type plugin int

func (p plugin) Example() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}

var Plugin plugin
