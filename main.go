package main

import (
	"fmt"
	"io/ioutil"
	"plugin"
	"strings"
)

type Plugin interface {
	Example()
}

func greatings() {
	fmt.Println()
	fmt.Println("Choose sample: ")

	files, err := ioutil.ReadDir("./plugins")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(" * ", strings.Split(file.Name(), ".")[0])
	}
	fmt.Println("To exit type: quit")
}

func executeExternalFunction(mod string) error {
	plug, err := plugin.Open("plugins/" + mod + ".so")
	if err != nil {
		return err
	}

	sym, err := plug.Lookup("Plugin")
	if err != nil {
		return err
	}

	var example Plugin
	example, ok := sym.(Plugin)

	if !ok {
		return fmt.Errorf("Unexpected type from module %s", mod)
	}

	example.Example()
	return nil
}

func main() {
	var sample string
	for sample != "quit" {
		greatings()
		fmt.Print("> ")
		fmt.Scanf("%s", &sample)
		if sample == "quit" {
			break
		}
		err := executeExternalFunction(sample)
		if err != nil {
			fmt.Println(err)
		}
	}
}
