package main

import "fmt"

type plugin int

func swap(x, y string) (string, string) {
	return y, x
}

func (p plugin) Example() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

var Plugin plugin
