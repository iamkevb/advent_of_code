package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input := readInput("./puzzle.input")
	// part1(input)
	part2(input)
}

func part1(input []string) {
	partNumbers := findPartNumber(input)
	sum := 0
	for _, n := range partNumbers {
		sum += n
	}
	fmt.Println("SUM: ", sum)
}

func part2(input []string) {
	gearRatios := findGearRatios(input)
	sum := 0
	for _, n := range gearRatios {
		sum += n
	}
	fmt.Println("SUM: ", sum)

}

func readInput(path string) []string {
	f, _ := os.Open(path)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

var reNumber = regexp.MustCompile(`\d+`)

func findPartNumber(lines []string) []int {
	partNumbers := []int{}
	for i, l := range lines {
		matches := reNumber.FindAllStringIndex(l, -1)
		for _, m := range matches {
			s, e := m[0], m[1]
			if isPartNumber(s, e, lines, i) {
				n, _ := strconv.Atoi(l[s:e])
				partNumbers = append(partNumbers, n)
			}
		}
	}
	return partNumbers
}

var reSymbol = regexp.MustCompile(`([^\d\.\n])`)

func isPartNumber(s int, e int, lines []string, l int) bool {
	start := max(0, s-1)
	end := min(e+1, max(len(lines[l])-1))
	if l > 0 {
		if reSymbol.Match([]byte(lines[l-1][start:end])) {
			return true
		}
	}
	if reSymbol.Match([]byte(lines[l][start:end])) {
		return true
	}
	if l < len(lines)-1 {
		if reSymbol.Match([]byte(lines[l+1][start:end])) {
			return true
		}
	}
	return false
}

var reStar = regexp.MustCompile(`\*`)

func findGearRatios(lines []string) []int {
	gearRatios := []int{}
	for i, l := range lines {
		matches := reStar.FindAllStringIndex(l, -1)
		for _, m := range matches {
			s := m[0]
			if ratio := isGear(s, lines, i); ratio > 0 {
				fmt.Println(i, ratio)
				gearRatios = append(gearRatios, ratio)
			}
		}
	}
	return gearRatios
}

func isGear(s int, lines []string, l int) int {
	var n1, n2 int = 0, 0
	if l > 0 {
		matches := reNumber.FindAllStringIndex(lines[l-1], -1)
		for _, m := range matches {
			if m[0] <= s+1 && m[1] >= s {
				n, _ := strconv.Atoi(lines[l-1][m[0]:m[1]])
				if n1 == 0 {
					n1 = n
				} else {
					n2 = n
				}
			}
		}
	}
	matches := reNumber.FindAllStringIndex(lines[l], -1)
	for _, m := range matches {
		if m[0] <= s+1 && m[1] >= s {
			n, _ := strconv.Atoi(lines[l][m[0]:m[1]])
			if n1 == 0 {
				n1 = n
			} else {
				n2 = n
			}
		}
	}
	if l < len(lines)-1 {
		matches := reNumber.FindAllStringIndex(lines[l+1], -1)
		for _, m := range matches {
			if m[0] <= s+1 && m[1] >= s {
				n, _ := strconv.Atoi(lines[l+1][m[0]:m[1]])
				if n1 == 0 {
					n1 = n
				} else {
					n2 = n
				}
			}
		}
	}
	return n1 * n2
}
