package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
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

	data, err := parse(input)
	if err != nil {
		log.Fatal(err)
	}

	printData(data)
}

func readInputFile(file string) (string, error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func parse(input string) (map[int][]int, error) {
	data := make(map[int][]int)

	for _, line := range strings.Split(input, "\n") {
		sections := strings.Split(line, ":")
		if len(sections) != 2 {
			return nil, fmt.Errorf("invalid input: %s", line)
		}
		sum, err := strconv.Atoi(sections[0])
		if err != nil {
			return nil, err
		}
		numbers := strings.Split(sections[1], " ")
		data[sum] = make([]int, len(numbers))
		for i, c := range numbers {
			if c == "" {
				continue
			}
			n, err := strconv.Atoi(c)
			if err != nil {
				return nil, err
			}
			data[sum][i] = n
		}
	}

	return data, nil
}

func printData(data map[int][]int) {
	for k, v := range data {
		fmt.Printf("%d: %v\n", k, v)
	}
}
