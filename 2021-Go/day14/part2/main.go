package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var polymer string
var rules map[string]rune = map[string]rune{}

var elements map[rune]int = map[rune]int{}

func main() {
	readInput("./input.txt")
	steps := 40

	var pairs map[string]int = makePairs(polymer)

	for i := 0; i < steps; i++ {
		pairs = growPairs(pairs)
	}

	//find max and min elements
	var min, max int = math.MaxInt, 0
	for _, c := range elements {
		if c < min {
			min = c
		}
		if c > max {
			max = c
		}
	}

	fmt.Printf("B %d, N %d, H %d, C %d\n", 'B', 'N', 'H', 'C')
	fmt.Println(elements)
	fmt.Printf("Max: %d, Min: %d, Diff: %d\n", max, min, max-min)
}

func growPairs(pairs map[string]int) map[string]int {
	result := make(map[string]int)

	for p, c := range pairs {
		m := rules[p]
		elements[m] += c
		// fmt.Printf("Adding %d to %s, total: %d. pair: %s\n", c, string(m), elements[m], p)
		p1 := string([]rune(p)[0]) + string(m)
		result[p1] += c
		p2 := string(m) + string([]rune(p)[1])
		result[p2] += c
	}

	return result
}

func makePairs(polymer string) map[string]int {
	pairs := make(map[string]int)
	runes := []rune(polymer)
	for i := 0; i < len(runes)-1; i++ {
		elements[runes[i]]++
		p := string(runes[i]) + string(runes[i+1])
		pairs[p]++
	}
	//increment last element in element map
	last := runes[len(runes)-1]
	elements[last]++
	return pairs
}

func readInput(path string) {
	f, _ := os.Open(path)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	polymer = scanner.Text()
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		var pl, pr, s2 string
		fmt.Sscanf(line, "%1s%1s -> %1s", &pl, &pr, &s2)
		rules[string([]rune(pl)[0])+string([]rune(pr)[0])] = []rune(s2)[0]
	}
}
