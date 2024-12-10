package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

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

func parseInput(input []string) ([]int, []int) {
	left := make([]int, 0)
	right := make([]int, 0)

	for _, val := range input {
		split := strings.Split(val, "   ")

		l, err := strconv.ParseInt(split[0], 10, 32);
		if err != nil {
			log.Fatalf("Failed to parse int: %v", err)
		}
		
		r, err := strconv.ParseInt(split[1], 10, 64)
		if err != nil {
			log.Fatalf("Failed to parse int: %v", err)
		}

		left = append(left, int(l))
		right = append(right, int(r))
	}

	return left, right
}

func diff(a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func main() {
	contents := loadInput("input.txt")
	left, right := parseInput(contents)
	total := 0

	sort.Ints(left)
	sort.Ints(right)

	for i := 0; i < len(left); i++ {
		total += diff(left[i], right[i])
	}

	log.Print(total)
}