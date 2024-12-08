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

	partOneOperators = []string{"+", "*"}
	partTwoOperators = []string{"+", "*", "|"}
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

	partOneSum, partTwoSum, err := compute(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part One: %d\n", partOneSum)
	fmt.Printf("Part Two: %d\n", partTwoSum)
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

func compute(data map[int][]int) (int, int, error) {
	partOneSum := 0
	partTwoSum := 0

	for k, v := range data {
		c, err := computeWithPerms(k, v, partOneOperators)
		if err != nil {
			return -1, -1, err
		}
		partOneSum += c
	}

	for k, v := range data {
		c, err := computeWithPerms(k, v, partTwoOperators)
		if err != nil {
			return -1, -1, err
		}
		partTwoSum += c
	}

	return partOneSum, partTwoSum, nil
}

func computeWithPerms(total int, numbers []int, operators []string) (int, error) {
	perms := generateOperatorPermutations(operators, len(numbers)-1)
	if *verbose {
		fmt.Println("--------------------")
		fmt.Printf("%d: %v\n", total, numbers)
		fmt.Printf("Perms: %v\n", perms)
	}
	for _, perm := range perms {
		parsedPerms := strings.Split(perm, "")
		if *verbose {
			fmt.Printf("Parsed perms: %v\n", parsedPerms)
		}
		sum := 0
		for i, num := range numbers {
			if num == -1 {
				continue
			}
			if sum == 0 {
				sum = num
				continue
			}
			switch parsedPerms[i-1] {
			case "+":
				sum += num
			case "*":
				sum *= num
			case "|":
				var err error
				sum, err = concatNumbers(sum, num)
				if err != nil {
					return -1, err
				}
			default:
				fmt.Printf("Invalid operator: %s\n", parsedPerms[i-1])
				continue
			}
		}

		if *verbose {
			fmt.Printf("Sum: %v\n", sum)
		}

		if sum == total {
			if *verbose {
				fmt.Printf("Sum %v matches with perm: %s\n", sum, perm)
			}
			return sum, nil
		}
	}
	return 0, nil
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

func concatNumbers(a, b int) (int, error) {
	aStr := strconv.Itoa(a)
	bStr := strconv.Itoa(b)

	aStr += bStr

	str, err := strconv.Atoi(aStr)
	if err != nil {
		return -1, err
	}

	return str, nil
}
