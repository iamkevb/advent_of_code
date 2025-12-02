package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const TEST = "./input-test.txt"
const DATA = "./input.txt"

type Range struct {
	l int64
	h int64
}

func main() {
	input := readInput(DATA)
	part1(input)
	part2(input)
}

func part1(input []Range) {
	var sum int64 = 0
	for _, r := range input {
		sum += sumInvalidIds(r)
	}
	fmt.Println("Part 1:", sum)
}

func sumInvalidIds(r Range) int64 {
	var sum int64 = 0
	for i := r.l; i <= r.h; i++ {
		s := strconv.FormatInt(i, 10)
		if isInvalidId(s) {
			sum += i
		}
	}
	return sum
}

func isInvalidId(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	n := len(s) / 2
	return s[:n] == s[n:]
}

func part2(input []Range) {
	var sum int64 = 0
	for _, r := range input {
		sum += sumInvalidIds2(r)
	}
	fmt.Println("Part 2:", sum)
}

func sumInvalidIds2(r Range) int64 {
	var sum int64 = 0
	for i := r.l; i <= r.h; i++ {
		s := strconv.FormatInt(i, 10)
		if isInvalidId2(s) {
			sum += i
		}
	}
	return sum
}

func isInvalidId2(s string) bool {
	sLen := len(s)

	// Start at half the length of the string, then check progressively smaller segments
	for rLen := sLen / 2; rLen > 0; rLen-- {
		// The repeating string must evenly divide into the string length
		if sLen%rLen != 0 {
			continue
		}

		repeats := true
		pattern := s[:rLen]

		// Check that substrings of segment len (rLen) from rLen on match the pattern.
		// Break out if they don't and try the next segment length.
		for i := rLen; i < sLen; i += rLen {
			if s[i:i+rLen] != pattern {
				repeats = false
				break
			}
		}

		// No need to check further segment lengths
		if repeats {
			return true
		}

	}
	return false
}

func readInput(path string) []Range {
	f, _ := os.ReadFile(path)
	ranges := []Range{}
	rangeStrs := strings.Split(string(f), ",")
	for _, v := range rangeStrs {
		limits := strings.Split(v, "-")
		low, _ := strconv.ParseInt(limits[0], 10, 64)
		high, _ := strconv.ParseInt(limits[1], 10, 64)
		r := Range{
			l: low,
			h: high,
		}
		ranges = append(ranges, r)
	}
	return ranges
}
