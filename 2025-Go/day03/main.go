package main

import (
	"bufio"
	"fmt"
	"os"
)

const TEST = "./input-test.txt"
const DATA = "./input.txt"

func main() {
	input := readInput(DATA)
	part1(input)
	part2(input)
}

func part1(input [][]int) {
	sum := int64(0)
	for _, b := range input {
		sum += maxJoltage(2, b)
	}
	fmt.Println("Part 1: ", sum)
}

func part2(input [][]int) {
	var sum int64 = 0
	for _, b := range input {
		sum += maxJoltage(12, b)
	}
	fmt.Println("Part 2: ", sum)
}

func maxJoltage(batteries int, bank []int) int64 {
	var sum int64 = 0
	idx := 0

	for i := 0; i < batteries; i++ {
		max := len(bank) - batteries + i + 1
		m, mi := indexOfMax(bank[idx:max])
		idx += mi + 1
		sum = sum*10 + int64(m)
	}

	return sum
}

func indexOfMax(a []int) (int, int) {
	max, idx := 0, 0
	for i, n := range a {
		if n > max {
			idx = i
			max = n
		}
	}
	return max, idx
}

func readInput(path string) [][]int {
	f, _ := os.Open(path)
	scanner := bufio.NewScanner(f)
	input := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		bank := []int{}
		for _, c := range line {
			bank = append(bank, int(c-'0'))
		}
		input = append(input, bank)
	}
	return input
}
