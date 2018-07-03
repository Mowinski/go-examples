package main

import "fmt"

type plugin int

func (p plugin) Example() {
	type Point struct {
		X, Y int
	}

	s := Point{42, 0}
	point := &s
	point.X = 0

	fmt.Println(s)
}

var Plugin plugin
