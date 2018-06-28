package main

import "fmt"

type plugin int

func (p plugin) Example() {
	people := "Guys"
	fmt.Println("Hello World")
	fmt.Printf("Hello %s\n", people)
	greatings := fmt.Sprintf("Hello %s", people)

	var age, name, surname string
	fmt.Print(greatings, ", what is your age: ")
	fmt.Scan(&age)
	fmt.Print("What is your name and surname: ")
	fmt.Scanf("%s %s", &name, &surname)

	if len(name) < 3 || len(surname) < 3 {
		err := fmt.Errorf("Wrong name")
		panic(err)
	}
}

var Plugin plugin
