package main

import (
	"io"
	"log"
	"os"
)

func main() {
	io.Copy(os.Stdout, readFromFile(os.Args[1]))
}

func readFromFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return file
}
