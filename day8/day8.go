package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
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

	antennaRegex = regexp.MustCompile(`[A-Za-z\d]`)
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

	data := parseMap(input)

	printMap(data)

	count, err := findAntinodes(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Antinode Count: %d\n", count)
}

func readInputFile(file string) (string, error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func parseMap(input string) [][]string {
	rows := strings.Split(input, "\n")
	m := make([][]string, len(rows))
	for i, row := range rows {
		m[i] = strings.Split(row, "")
	}
	return m
}

func printMap(m [][]string) {
	for _, row := range m {
		fmt.Println(strings.Join(row, " "))
	}
	fmt.Println()
}

func findAntinodes(m [][]string) (int, error) {
	antinodes := 0
	for i, row := range m {
		for j, col := range row {
			if antennaRegex.MatchString(col) {
				fmt.Printf("Row: %d, Col: %d, Value: %s\n", i, j, col)
			}
		}
	}
	return antinodes, nil
}
