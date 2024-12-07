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
	verbose     = flag.Bool("v", false, "Enable verbose output")
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

	m, startingCol, startingRow, err := parseMap(input)
	if err != nil {
		log.Fatal(err)
	}

	positions := traverse(m, startingCol, startingRow)

	fmt.Printf("Unique Positions: %d", positions)
}

func readInputFile(file string) (string, error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func parseMap(input string) ([][]string, int, int, error) {
	rows := strings.Split(input, "\n")
	m := make([][]string, len(rows))
	startingCol := 0
	startingRow := 0
	for i, row := range rows {
		cols := strings.Split(row, "")
		m[i] = cols
		for j, col := range cols {
			if col == "^" {
				startingCol = j
				startingRow = i
			}
		}
	}
	return m, startingCol, startingRow, nil
}

func printMap(m [][]string) {
	for _, row := range m {
		fmt.Println(strings.Join(row, " "))
	}
}

func traverse(m [][]string, col, row int) int {
	uniquePositions := 0

	printMap(m)

	return uniquePositions
}
