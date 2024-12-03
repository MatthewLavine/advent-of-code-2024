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
	demo        = flag.Bool("demo", false, "Use demo input")
	enablePprof = flag.Bool("pprof", false, "Enable pprof")

	instructionRegex = regexp.MustCompile(`mul\(\d+,\d+\)`)
)

type Instruction struct {
	operation string
	a, b      int
}

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

	instructions, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(instructions)

	val, err := compute(instructions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(val)
}

func readInputFile(file string) (string, error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func parseInput(input string) ([]Instruction, error) {
	var instructions []Instruction
	matches := instructionRegex.FindAllString(input, -1)
	if matches == nil {
		return nil, fmt.Errorf("no instructions found")
	}
	for _, match := range matches {
		op := match[0:3]
		args := strings.Split(match[4:len(match)-1], ",")
		a, err := strconv.Atoi(args[0])
		if err != nil {
			return nil, err
		}
		b, err := strconv.Atoi(args[1])
		if err != nil {
			return nil, err
		}
		instructions = append(instructions, Instruction{
			operation: op,
			a:         a,
			b:         b,
		})
	}
	return instructions, nil
}

func compute(instructions []Instruction) (int, error) {
	ret := 0
	for _, instruction := range instructions {
		switch instruction.operation {
		case "mul":
			ret += instruction.a * instruction.b
		}
	}
	return ret, nil
}
