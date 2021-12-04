package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := parseInput("./input.txt")
	o2 := computeO2(input)
	co2 := computeCO2(input)

	fmt.Printf("Oxygen generator: %d, CO2 Scrubber: %d, Life Support Rating: %d\n", o2, co2, o2*co2)
}

func computeCO2(input []string) int64 {
	start := ""
	s := input

	for i := 0; i < len(input[0]); i++ {
		_, least := mostLeast(s, i)
		start += least

		remaining := []string{}
		for _, v := range s {
			if strings.HasPrefix(v, start) {
				remaining = append(remaining, v)
			}
		}

		s = remaining

		if len(s) == 1 {
			break
		}
	}
	val, _ := strconv.ParseInt(s[0], 2, 64)
	return val
}

func computeO2(input []string) int64 {
	start := ""
	s := input

	for i := 0; i < len(input[0]); i++ {
		most, _ := mostLeast(s, i)
		start += most

		remaining := []string{}
		for _, v := range s {
			if strings.HasPrefix(v, start) {
				remaining = append(remaining, v)
			}
		}

		s = remaining

		if len(s) == 1 {
			break
		}
	}
	val, _ := strconv.ParseInt(s[0], 2, 64)
	return val
}

func mostLeast(data []string, column int) (string, string) {
	oneCount := 0

	for _, v := range data {
		if v[column:column+1] == "1" {
			oneCount += 1
		}
	}

	if oneCount*2 >= len(data) {
		return "1", "0"
	}
	return "0", "1"
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
