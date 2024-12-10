package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func Map[T any, U any](input []T, fn func(T) U) []U {
	results := make([]U, len(input))
	for i, v := range input {
		results[i] = fn(v)
	}

	return results
}

func loadInput(file string) []string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("Error opening file: %v %v", file, err)
	}
	defer f.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v %v", file, err)
	}

	return lines
}

func parseInput(input []string) [][]int {
	parsed := make([][]int, 0)
	for _, val := range input {
		numStrs := strings.Split(val, " ")
		nums := Map(numStrs, func(str string) int {
			n, err := strconv.ParseInt(str, 10, 32)
			if err != nil {
				log.Fatalf("Unable to parse string to int: %v", err)
			}

			return int(n)
		})

		parsed = append(parsed, nums)
	}

	return parsed
}

func diff(a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func validDiff(diff int) bool {
	abs := int(math.Abs(float64(diff)))
	log.Print(abs)
	return abs >= 1 && abs <= 3
}

func safe(report []int) bool {
	// Need to obtain the intial direction (increasing or decreasing)
	initial := report[1] - report[0]
	if !validDiff(initial) {
		return false
	}

	increasing := initial > 0
	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]
		if !validDiff(diff) {
			return false
		}

		if (increasing && diff < 0) || (!increasing && diff > 0) {
			return false
		}
	}

	return true
}

func main() {
	input := loadInput("input.txt")
	parsedInput := parseInput(input);
	total := 0

	for _, report := range parsedInput {
		if safe(report) {
			total += 1
		}
	}

	log.Print(total)
}