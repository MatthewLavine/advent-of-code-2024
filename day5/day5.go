package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"slices"
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

type rule struct {
	l, r int
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

	rules, updates, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	sum, err := computePart1(rules, updates)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part1 sum: %d\n", sum)

	sum, err = computePart2(rules, updates)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part2 sum: %d\n", sum)
}

func readInputFile(file string) (string, error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func parseInput(input string) ([]rule, [][]int, error) {
	var rules []rule
	var updates [][]int
	for _, line := range strings.Split(input, "\n") {
		if strings.Contains(line, "|") {
			rawRule := strings.Split(line, "|")
			l, err := strconv.Atoi(rawRule[0])
			if err != nil {
				return nil, nil, err
			}
			r, err := strconv.Atoi(rawRule[1])
			if err != nil {
				return nil, nil, err
			}
			rule := rule{
				l: l,
				r: r,
			}
			rules = append(rules, rule)
		}
		if strings.Contains(line, ",") {
			newUpdates, err := stringListToIntList(strings.Split(line, ","))
			if err != nil {
				return nil, nil, err
			}
			updates = append(updates, newUpdates)
		}
		continue
	}
	return rules, updates, nil
}

func stringListToIntList(strList []string) ([]int, error) {
	var intList []int
	for _, str := range strList {
		i, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		intList = append(intList, i)
	}
	return intList, nil
}

func computePart1(rules []rule, updates [][]int) (int, error) {
	var sum int
nextUpdate:
	for _, update := range updates {
		satisfiesRules := true
		for _, rule := range rules {
			lValid := slices.Index(update, rule.l)
			rValid := slices.Index(update, rule.r)
			if lValid == -1 || rValid == -1 {
				continue
			}
			if lValid < rValid {
				if *verbose {
					fmt.Printf("Update %v satisfies rule %v\n", update, rule)
				}
			} else {
				satisfiesRules = false
				if *verbose {
					fmt.Printf("Update %v does not satisfy rule %v\n", update, rule)
				}
				continue nextUpdate
			}
		}
		if satisfiesRules {
			sum += update[len(update)/2]
		}
	}
	return sum, nil
}

func computePart2(rules []rule, updates [][]int) (int, error) {
	var badUpdates [][]int
nextUpdate:
	for _, update := range updates {
		for _, rule := range rules {
			lValid := slices.Index(update, rule.l)
			rValid := slices.Index(update, rule.r)
			if lValid == -1 || rValid == -1 {
				continue
			}
			if lValid >= rValid {
				// Collect bad updates.
				badUpdates = append(badUpdates, update)
				continue nextUpdate
			}
		}
	}
	var sum int
	for _, update := range badUpdates {
		satisfiesRules := true
		for _, rule := range rules {
			lValid := slices.Index(update, rule.l)
			rValid := slices.Index(update, rule.r)
			if lValid == -1 || rValid == -1 {
				continue
			}
			if lValid >= rValid {
				// Make update conform to rule.
				update[lValid], update[rValid] = update[rValid], update[lValid]
			}

		}
		if satisfiesRules {
			sum += update[len(update)/2]
		}
	}

	return computePart1(rules, badUpdates)
}
