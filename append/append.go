package main

import "fmt"

type plugin int

func (p plugin) Example() {
	var num []int
	num = append(num, 42)
	fmt.Println(num)
}

var Plugin plugin
