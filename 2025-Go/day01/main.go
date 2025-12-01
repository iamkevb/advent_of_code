package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const TEST = "./input-test.txt"
const DATA = "./input.txt"
const RAUL = "./raul.txt"

func main() {
	input := readInput(DATA)
	count := part1(input, 50)
	fmt.Println("Part 1: ", count)
	count = part2(input, 50)
	fmt.Println("Part 2: ", count)
}

func part1(input []int, start int) int {
	pos := start
	count := 0
	for _, val := range input {
		pos += val

		for pos >= 100 {
			pos = pos % 100
		}
		for pos < 0 {
			pos = 100 + pos
		}
		if pos == 0 {
			count++
		}
	}
	return count
}

func part2(input []int, start int) int {
	pos := start
	count := 0

	for _, val := range input {
		revs := val / 100
		if revs < 0 {
			revs = -revs
		}
		rem := val % 100

		count += revs
		newPos := pos + rem

		if pos != 0 {
			if newPos <= 0 {
				count++
			} else if newPos >= 100 {
				count++
			}
		}

		if newPos > 99 {
			newPos -= 100
		}
		if newPos < 0 {
			newPos += 100
		}

		pos = newPos
	}

	return count
}

func readInput(path string) []int {
	f, _ := os.Open(path)
	scanner := bufio.NewScanner(f)
	rotations := []int{}
	reDir := regexp.MustCompile(`[LR]`)
	reVal := regexp.MustCompile(`\d*$`)
	for scanner.Scan() {
		line := scanner.Text()
		dir := reDir.Find([]byte(line))
		val, _ := strconv.Atoi(string(reVal.Find([]byte(line))))
		if string(dir) == string('L') {
			val *= -1
		}

		rotations = append(rotations, val)
	}
	return rotations
}
