package main

import (
	"fmt"
	"io"
)

type plugin int

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(p []byte) (int, error) {
	fmt.Println(p)
	return len(p), nil
}

func NeedWriter(w io.Writer, text string) {
	w.Write([]byte(text))
}

func (p plugin) Example() {
	cw := ConsoleWriter{}
	NeedWriter(cw, "Hello World")
}

var Plugin plugin
