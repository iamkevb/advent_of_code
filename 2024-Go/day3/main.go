package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Part1()
	Part2()
}

func readInput(path string) string {
	f, _ := os.ReadFile(path)
	return string(f)
}

// PART1
func Part1() {
	total := 0

	// s := readInput("./input/test.txt")
	s := readInput("./input/input.txt")
	regex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := regex.FindAllStringSubmatch(s, -1)

	for _, m := range matches {
		m1, _ := strconv.Atoi(m[1])
		m2, _ := strconv.Atoi(m[2])
		total += m1 * m2
	}
	fmt.Println(total)
}

// Part 2

func Part2() {
	total := 0

	// s := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	s := readInput("./input/input.txt")

	// Remove new lines - so I don't have to account for don't()xxx\n not matching
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, "\r", "")

	// Remove don't sections - .+? is non-greedy and ensures the next do() is captured
	regex := regexp.MustCompile(`don't\(\).+?do\(\)`)
	s = regex.ReplaceAllString(s, "")
	fmt.Println(s)

	regex = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := regex.FindAllStringSubmatch(s, -1)

	for _, m := range matches {
		m1, _ := strconv.Atoi(m[1])
		m2, _ := strconv.Atoi(m[2])
		total += m1 * m2
	}
	fmt.Println(total)
}
