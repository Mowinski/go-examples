package main

import "fmt"

type plugin int

func count() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

func (p plugin) Example() {
	next := count()

	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
}

var Plugin plugin
