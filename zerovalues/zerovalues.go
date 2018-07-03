package main

import "fmt"

type plugin int

var (
	b bool
	i int
	f float32
	c complex64
	s string
)

func (p plugin) Example() {
	fmt.Println(b, i, f, c)
	fmt.Printf("%q\n", s)
}

var Plugin plugin
