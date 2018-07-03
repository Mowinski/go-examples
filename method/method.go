package main

import (
	"fmt"
	"math"
)

type plugin int

type Vector struct {
	X, Y float64
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (p plugin) Example() {
	v := Vector{3, 4}
	fmt.Println(v.Length())
}

var Plugin plugin
