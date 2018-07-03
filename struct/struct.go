package main

import "fmt"

type plugin int

func (p plugin) Example() {
	type Point struct {
		X int
		Y int
	}

	p1 := Point{1, 2}
	fmt.Println(p1)
	fmt.Println(p1.Y)

	p2 := Point{X: 2}
	p3 := Point{}

	fmt.Println(p2, p3)
}

var Plugin plugin
