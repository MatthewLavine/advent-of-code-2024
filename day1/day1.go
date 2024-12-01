package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"sort"
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
)

func main() {
	flag.Parse()

	path := inputFile
	if *demo {
		path = demoInputFile
	}

	input, err := readInputFile(path)
	if err != nil {
		log.Fatal(err)
	}

	left, right, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	distance, err := calculateListDistance(left, right)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Distance: %d\n", distance)

	similar, err := calculateListSimilarity(left, right)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Similarity: %d\n", similar)
}

func readInputFile(file string) (string, error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func parseInput(input string) ([]int, []int, error) {
	var left []int
	var right []int
	for i, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		var found []string

		for _, char := range strings.Fields(line) {

			if numberRegex.MatchString(char) {
				found = append(found, char)
			}
		}

		if len(found) != 2 {
			return nil, nil, fmt.Errorf("not enough numbers found in line %d: %s\n found %s", i, line, found)
		}

		l, err := strconv.Atoi(found[0])
		if err != nil {
			return nil, nil, err
		}
		r, err := strconv.Atoi(found[1])
		if err != nil {
			return nil, nil, err
		}

		left = append(left, l)
		right = append(right, r)
	}
	return left, right, nil
}

func calculateListDistance(left, right []int) (int, error) {
	if len(left) != len(right) {
		return -1, fmt.Errorf("left and right lists are not the same length")
	}

	diffs := make([]int, len(left))

	for i := 0; i < len(left); i++ {
		l, err := nthSmallest(i, left)
		if err != nil {
			return -1, err
		}
		r, err := nthSmallest(i, right)
		if err != nil {
			return -1, err
		}
		diff := r - l
		if diff < 0 {
			diff = -diff
		}
		diffs[i] = diff
	}

	distance := 0
	for _, diff := range diffs {
		distance += diff
	}

	return distance, nil
}

func nthSmallest(n int, arr []int) (int, error) {
	if n > len(arr) {
		return -1, fmt.Errorf("n is greater than the length of the array")
	}
	copy := slices.Clone(arr)
	sort.Ints(copy)
	return copy[n], nil
}

func calculateListSimilarity(left, right []int) (int, error) {
	if len(left) != len(right) {
		return -1, fmt.Errorf("left and right lists are not the same length")
	}

	similarities := make([]int, len(left))

	for i := 0; i < len(left); i++ {
		similarities[i] = left[i] * count(left[i], right)
	}

	sum := 0
	for _, similarity := range similarities {
		sum += similarity
	}

	return sum, nil
}

func count(n int, arr []int) int {
	count := 0
	for _, i := range arr {
		if i == n {
			count++
		}
	}
	return count
}
