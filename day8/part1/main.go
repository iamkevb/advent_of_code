package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseOutputDigits(text string) []int {
	split := strings.Split(text, "|")
	outputText := split[1]
	output := strings.Split(outputText, " ")
	var result []int
	for _, v := range output {
		switch len(v) {
		case 2:
			result = append(result, 1)
		case 3:
			result = append(result, 7)
		case 4:
			result = append(result, 4)
		case 7:
			result = append(result, 8)
		}
	}
	return result
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic("Cannot open file")
	}
	defer f.Close()

	count := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		digits := parseOutputDigits(line)
		for _, d := range digits {
			switch d {
			case 1, 4, 7, 8:
				count += 1
			}
		}
	}
	fmt.Println(count)
}
