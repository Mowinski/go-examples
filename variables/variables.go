package main

import "fmt"

type plugin int

var foo int

func (p plugin) Example() {
	foo = 42

	var bar int = 1337
	baz := "Hello"
	fmt.Println(foo, bar, baz)

	foo, bar = 0, -1
	fmt.Println(foo, bar)

	const Pi = 3.1415
	fmt.Println(Pi)
}

var Plugin plugin
