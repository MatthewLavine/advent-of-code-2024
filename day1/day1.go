package main

import (
	"fmt"
	"log"
	"os"
)

const inputFile = "demo.txt"

func main() {
	input := readInputFile()
	fmt.Printf("%s\n", input)
}

func readInputFile() string {
	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}
