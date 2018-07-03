package main

import "fmt"

type plugin int

func (p plugin) Example() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	foo := 1

	for ; foo < 42 ; {
		foo += foo
	}

	fmt.Println(foo)

	foo = 1
	for foo < 42 {
		foo += foo
	}

	fmt.Println(foo)
}

var Plugin plugin
