package main

import "fmt"

type plugin int

func (p plugin) Example() {
	foo := 42

	pointer := &foo
	fmt.Println(*pointer)
	fmt.Println(pointer)

	*pointer = 1
	fmt.Println(foo)

}

var Plugin plugin
