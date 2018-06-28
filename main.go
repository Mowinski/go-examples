package main

import (
	"fmt"
	"plugin"
)

type Plugin interface {
	Example()
}

var sampleList = [...]string{
	"print",
	"bytes",
	"writer",
	"json",
}

func greatings() {
	fmt.Println()
	fmt.Println("Choose sample: ")
	for _, app := range sampleList {
		fmt.Println(" * ", app)
	}
	fmt.Println("To exit type: quit")
}

func executeExternalFunction(mod string) error {
	plug, err := plugin.Open("build/" + mod + ".so")
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
