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

	processReports(reports)
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

func processReports(reports [][]int) {
	rawSafeReports := 0
	dampendedSafeReports := 0
	for _, report := range reports {
		if isReportSafe(report, false) {
			rawSafeReports++
		}
		if isReportSafe(report, true) {
			dampendedSafeReports++
		}
	}
	fmt.Printf("Safe reports: %d\n", rawSafeReports)
	fmt.Printf("Dampened safe reports: %d\n", dampendedSafeReports)
}

func isReportSafe(report []int, dampener bool) bool {
	first := true
	second := false
	prev := 0
	increasing := false
	decreasing := false
	ignored := false
	for _, level := range report {
		if first {
			first = false
			second = true
			prev = level
			continue
		}
		if second {
			second = false
			if level > prev {
				increasing = true
			} else if level < prev {
				decreasing = true
			}
		}
		if increasing && level < prev {
			if dampener {
				if ignored {
					return false
				}
				ignored = true
				continue
			}
			return false
		}
		if decreasing && level > prev {
			if dampener {
				if ignored {
					return false
				}
				ignored = true
				continue
			}
			return false
		}
		diff := diff(level, prev)
		if diff < 1 || diff > 3 {
			if dampener {
				if ignored {
					return false
				}
				ignored = true
				continue
			}
			return false
		}
		prev = level
	}
	return true
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
