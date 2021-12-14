package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var polymer string
var insertionPairs map[string]string = map[string]string{}

func main() {
	readInput("./input.txt")
	steps := 10
	for i := 0; i < steps; i++ {
		polymer = expandPolymer()
	}
	// fmt.Println(polymer)
	fmt.Println(doMath())
}

func doMath() int {
	elements := make(map[rune]int)
	for _, v := range polymer {
		elements[v]++
	}
	var min, max int = math.MaxInt, 0
	for _, v := range elements {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return max - min
}

func expandPolymer() string {
	expanded := string(polymer[0])
	for i := 0; i < len(polymer)-1; i++ {
		rune1 := polymer[i]
		rune2 := polymer[i+1]
		key := string(rune1) + string(rune2)
		expanded += insertionPairs[key] + string(rune2)
	}
	return expanded
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
		var s1, s2 string
		fmt.Sscanf(line, "%2s -> %s", &s1, &s2)
		insertionPairs[s1] = s2
	}
}
