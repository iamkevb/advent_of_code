package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const TEST = "./input-test.txt"
const DATA = "./input.txt"

func main() {
	input := DATA
	worksheet := readInput(input)
	part1(worksheet)

	w2 := readCephalopod(input)
	part2(w2)
}

func part1(worksheet [][]string) {
	answers := []int64{}

	for _, row := range worksheet {
		values := []int64{}
		for _, v := range row {
			switch v {
			case "+":
				answers = append(answers, sum(values))
			case "*":
				answers = append(answers, product(values))
			default:
				iVal, e := strconv.ParseInt(v, 10, 64)
				if e != nil {
					panic(e.Error())
				}
				values = append(values, iVal)
			}
		}
	}

	fmt.Println("Part 1:", sum(answers))
}

func sum(values []int64) int64 {
	s := int64(0)
	for _, v := range values {
		s += v
	}
	return s
}

func product(values []int64) int64 {
	s := int64(1)
	for _, v := range values {
		s *= v
	}
	return s
}

func readInput(path string) [][]string {
	f, _ := os.ReadFile(path)
	scanner := bufio.NewScanner(bytes.NewReader(f))
	lines := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, strings.Fields(line))
	}
	return transpose(lines)
}

func transpose(s [][]string) [][]string {
	t := [][]string{}
	for r := 0; r < len(s); r++ {
		for c := 0; c < len(s[r]); c++ {
			if len(t) <= c {
				t = append(t, []string{s[r][c]})
			} else {
				t[c] = append(t[c], s[r][c])
			}
		}
	}
	return t
}

func part2(worksheet [][]rune) {
	answers := []int64{}
	rows := len(worksheet)
	cols := len(worksheet[0])

	values := []int64{}
	operation := '+'

	for c := cols - 1; c >= 0; c-- {
		value := int64(0)
		for r := 0; r < rows; r++ {
			switch worksheet[r][c] {
			case '+':
				operation = worksheet[r][c]
			case '*':
				operation = worksheet[r][c]
			case ' ':
				continue
			default:
				value = value*10 + int64(worksheet[r][c]-'0')
			}
		}
		if value != 0 {
			values = append(values, value)
		}
		if value == 0 || c == 0 {
			switch operation {
			case '+':
				answers = append(answers, sum(values))
			case '*':
				answers = append(answers, product(values))
			}
			values = []int64{}
		}
	}

	fmt.Println("Part 2:", sum(answers))
}

func readCephalopod(path string) [][]rune {
	f, _ := os.ReadFile(path)
	scanner := bufio.NewScanner(bytes.NewReader(f))
	input := [][]rune{}
	max := -1
	for scanner.Scan() {
		runes := []rune(scanner.Text())
		if len(runes) != max && max != -1 {
			panic("rows are different lengths")
		}
		max = len(runes)
		input = append(input, []rune(scanner.Text()))
	}
	return input
}
