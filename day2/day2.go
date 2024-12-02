package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime/pprof"
	"strconv"
	"strings"
)

const (
	demoInputFile = "demo.txt"
	inputFile     = "input.txt"
)

var (
	numberRegex = regexp.MustCompile("^[0-9]*$")
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

	reports, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data: %v\n", reports)
}

func readInputFile(file string) (string, error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func parseInput(input string) ([][]int, error) {
	var reports [][]int
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		var report []int
		for _, char := range strings.Fields(line) {
			level, err := strconv.Atoi(char)
			if err != nil {
				return nil, err
			}
			report = append(report, level)
		}
		reports = append(reports, report)
	}
	return reports, nil
}
