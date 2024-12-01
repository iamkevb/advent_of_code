package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	part1()
	part2()
}

// PART1
func part1() {
	// loc1, loc2 := readInput("./input/test.txt")
	loc1, loc2 := readInput("./input/input.txt")
	loc1 = sortLocations(loc1)
	loc2 = sortLocations(loc2)
	total := computeDistances(loc1, loc2)
	fmt.Println(total)
}

func readInput(path string) ([]int, []int) {
	f, _ := os.Open(path)
	scanner := bufio.NewScanner(f)
	loc1 := []int{}
	loc2 := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		var l1, l2 int
		fmt.Sscanf(line, "%d %d", &l1, &l2)
		loc1 = append(loc1, l1)
		loc2 = append(loc2, l2)
	}
	return loc1, loc2
}

func sortLocations(locs []int) []int {
	sort.Slice(locs, func(i, j int) bool {
		return locs[i] < locs[j]
	})
	return locs
}

func computeDistances(loc1 []int, loc2 []int) int {
	var total int = 0
	for i := 0; i < len(loc1); i++ {
		dist := loc1[i] - loc2[i]
		if dist < 0 {
			dist = -dist
		}
		total += dist
	}
	return total
}

// Part 2

func part2() {
	// loc1, loc2 := readInput("./input/test.txt")
	loc1, loc2 := readInput("./input/input.txt")

	loc2 = sortLocations(loc2)
	occurances := make(map[int]int)
	for _, val := range loc2 {
		occurances[val] = occurances[val] + 1
	}

	similarity := 0
	for _, val := range loc1 {
		similarity += val * occurances[val]
	}
	fmt.Println(similarity)
}
