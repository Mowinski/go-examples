package main

import (
	"bytes"
	"encoding/base64"
	"io"
	"os"
)

type plugin int

func (p plugin) Example() {
	var alphabet bytes.Buffer

	alphabet.Write([]byte("abcdefghijklmnoprstuwxyz\n"))
	alphabet.WriteTo(os.Stdout)

	buf := bytes.NewBufferString("R29waGVycyBydWxlIQ==")
	dec := base64.NewDecoder(base64.StdEncoding, buf)
	io.Copy(os.Stdout, dec)
}

var Plugin plugin
