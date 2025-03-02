package main

import (
	"io"
	"os"
)

func main() {
	args := os.Args
	txt_file := args[1]
	file, _ := os.Open(txt_file)
	io.Copy(os.Stdout, file)
}
