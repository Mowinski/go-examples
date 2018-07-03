package main

import "fmt"

type plugin int

func (p plugin) Example() {
	primes := []int{1, 2, 3, 5, 7, 11}

	fmt.Println(primes[1:3])
	fmt.Println(primes[3])
	fmt.Println(primes[4:])
	fmt.Println(primes[:2])
	fmt.Println(primes[:])
}

var Plugin plugin
