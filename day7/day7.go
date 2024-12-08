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

	sum, err := compute(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Sum: %d\n", sum)
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
		convertedNumbers := make([]int, 0)
		for _, c := range numbers {
			if c == "" {
				continue
			}
			n, err := strconv.Atoi(c)
			if err != nil {
				return nil, err
			}
			convertedNumbers = append(convertedNumbers, n)
		}
		data[sum] = convertedNumbers
	}

	return data, nil
}

func compute(data map[int][]int) (int, error) {
	sum := 0

	for k, v := range data {
		sum += findOperatorsForTotal(k, v)
	}

	return sum, nil
}

func findOperatorsForTotal(total int, numbers []int) int {
	if *verbose {
		fmt.Println("--------------------")
		fmt.Printf("%d: %v\n", total, numbers)
	}

	operators := []string{"+", "*"}

	perms := generateOperatorPermutations(operators, len(numbers)-1)

	if *verbose {
		fmt.Printf("Perms: %v\n", perms)
	}

	for _, perm := range perms {
		parsedPerms := strings.Split(perm, "")
		sum := 0
		for i, num := range numbers {
			if i == 0 {
				sum = num
				continue
			}
			switch parsedPerms[i-1] {
			case "+":
				sum += num
			case "*":
				sum *= num
			}
		}

		if sum == total {
			if *verbose {
				fmt.Printf("Found: %s\n", perm)
			}
			return sum
		}
	}

	return 0
}

func generateOperatorPermutations(operators []string, length int) []string {
	if length == 1 {
		return operators
	}

	perms := make([]string, 0)
	for _, op := range operators {
		for _, p := range generateOperatorPermutations(operators, length-1) {
			perms = append(perms, op+p)
		}
	}

	return perms
}
