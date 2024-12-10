package main

import (
	"bufio"
	"log"
	"os"
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

func main() {
	contents := loadInput("input.txt")
	left, right := parseInput(contents)
	sim := make(map[int]int)
	total := 0

	for _, num := range left {
		sim[num] = 0
	}
	for _, num := range right {
		if _, ok := sim[num]; ok {
			sim[num] += 1
		}
	}
	for key := range sim {
		total += sim[key] * key
	}

	log.Print(total)
}