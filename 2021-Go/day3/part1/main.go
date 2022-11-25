package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	input := parseInput("./input.txt")
	computeGamma(input)
}

func computeGamma(input []string) {
	gamma := 0
	epsilon := 0

	lineLen := utf8.RuneCountInString(input[0]) //assume lines are same length
	columns := make([]string, lineLen)

	for _, v := range input {
		runes := []rune(v)
		for i, r := range runes {
			columns[i] += string(r) // builder would be better.
		}
	}

	for i, col := range columns {
		most, least := mostLeast(col)
		gamma += most << (lineLen - 1 - i)
		epsilon += least << (lineLen - 1 - i)
	}

	fmt.Printf("gamma: %d, epsilon: %d, product: %d\n", gamma, epsilon, gamma*epsilon)
}

func mostLeast(column string) (int, int) {
	oneCount := 0

	for _, v := range column {
		if v == '1' {
			oneCount += 1
		}
	}

	if oneCount*2 > utf8.RuneCountInString(column) {
		return 1, 0
	}
	return 0, 1
}

func parseInput(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		panic("Failed to open input")
	}

	var input []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		input = append(input, text)
	}
	return input
}
