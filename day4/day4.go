package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strings"
)

const (
	demoInputFile = "demo.txt"
	inputFile     = "input.txt"
)

var (
	demo        = flag.Bool("demo", false, "Use demo input")
	enablePprof = flag.Bool("pprof", false, "Enable pprof")
)

func main() {
	flag.Parse()

	if *enablePprof {
		f, err := os.Create("profile.prof")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			panic(err)
		}
		defer pprof.StopCPUProfile()
	}

	path := inputFile
	if *demo {
		path = demoInputFile
	}

	input, err := readInputFile(path)
	if err != nil {
		log.Fatal(err)
	}

	matrix, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	processMatrix(matrix)
}

func readInputFile(file string) (string, error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func parseInput(input string) ([][]string, error) {
	var matrix [][]string
	for _, line := range strings.Split(input, "\n") {
		matrix = append(matrix, strings.Split(line, ""))
	}
	return matrix, nil
}

func processMatrix(matrix [][]string) {
	for _, row := range matrix {
		for _, col := range row {
			fmt.Printf("%s ", col)
		}
		fmt.Println()
	}
}
