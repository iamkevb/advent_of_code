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

type Range struct {
	l, h int64
}

func main() {
	ranges, ingredients := readInput(DATA)
	part1(ranges, ingredients)
	part2(ranges)
}

func part1(ranges []Range, ingredients []int64) {
	count := 0
	for _, v := range ingredients {
		if isFresh(ranges, v) {
			count++
		}
	}
	fmt.Println("Part 1: ", count)
}

func isFresh(ranges []Range, ingredient int64) bool {
	for _, r := range ranges {
		if ingredient >= r.l && ingredient <= r.h {
			return true
		}
	}
	return false
}

func part2(ranges []Range) {
	combinedRanges := make(map[string]Range)
	rangeCount := len(ranges)

	for {
		for _, r := range ranges {
			combinedRanges = combine(r, combinedRanges)
		}
		if rangeCount == len(combinedRanges) {
			break
		}
		rangeCount = len(combinedRanges)
	}

	count := int64(0)
	for _, v := range combinedRanges {
		count += v.h - v.l + 1
	}
	fmt.Println("Part 2: ", count)
}

func combine(r Range, combinedRanges map[string]Range) map[string]Range {
	combined := false
	for k, v := range combinedRanges {
		rCombined := v.combine(r)
		if len(rCombined) == 1 {
			delete(combinedRanges, r.key())
			delete(combinedRanges, k)
			combinedRanges[v.key()] = rCombined[0]
			combined = true
		}
	}
	if !combined {
		combinedRanges[r.key()] = r
	}
	return combinedRanges
}

func (r Range) key() string {
	return fmt.Sprintf("%d-%d", r.l, r.h)
}

func (r Range) combine(r2 Range) []Range {
	if r2.h < r.l || r2.l > r.h {
		return []Range{r, r2}
	}
	if r2.h == r.h && r2.l == r.l {
		return []Range{r}
	}
	return []Range{
		{
			l: min(r.l, r2.l),
			h: max(r.h, r2.h),
		},
	}
}

func min(i1, i2 int64) int64 {
	if i1 < i2 {
		return i1
	}
	return i2
}

func max(i1, i2 int64) int64 {
	if i1 > i2 {
		return i1
	}
	return i2
}

func readInput(path string) ([]Range, []int64) {
	f, _ := os.ReadFile(path)
	scanner := bufio.NewScanner(bytes.NewReader(f))
	ranges := []Range{}
	ingredients := []int64{}
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, "-")
		if len(values) == 2 {
			l, _ := strconv.ParseInt(values[0], 10, 64)
			h, _ := strconv.ParseInt(values[1], 10, 64)
			ranges = append(ranges, Range{l, h})
		}
		if len(values) == 1 && len(values[0]) > 0 {
			i, _ := strconv.ParseInt(values[0], 10, 64)
			ingredients = append(ingredients, i)
		}
	}
	return ranges, ingredients
}
