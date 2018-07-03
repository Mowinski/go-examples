package main

import "fmt"

type plugin int

func sum(foo, bar int) int {
	return foo + bar
}

func (p plugin) Example() {
	fmt.Println(sum(1, 42))
}

var Plugin plugin
