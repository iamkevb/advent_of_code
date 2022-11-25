package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type NodeMap map[string][]string

var caves NodeMap = NodeMap{}

func main() {
	createNodeMap("./input.txt")

	paths := findAllPaths()
	fmt.Println(len(paths))
}

func findAllPaths() [][]string {
	path := []string{"start"}
	return findPaths(path, "end")
}

func findPaths(pathStart []string, end string) [][]string {
	start := pathStart[len(pathStart)-1]
	if start == end {
		return [][]string{pathStart}
	}
	if contains(pathStart, "end") {
		fmt.Println("shouldn't happen")
	}
	node := caves[start]
	completePaths := [][]string{}
	for _, n := range node {
		if strings.ToLower(n) == n && contains(pathStart, n) {
			continue //invalid step
		}
		ps := append(pathStart, n)
		p := findPaths(ps, end)
		completePaths = append(completePaths, p...)
	}

	return completePaths
}

func contains(arr []string, value string) bool {
	for _, s := range arr {
		if s == value {
			return true
		}
	}
	return false
}

func createNodeMap(path string) {
	f, _ := os.Open(path)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "-")
		c1 := split[0]
		c2 := split[1]

		cave1 := caves[c1]
		if cave1 == nil {
			caves[c1] = []string{c2}
		} else {
			caves[c1] = append(caves[c1], c2)
		}

		cave2 := caves[c2]
		if cave2 == nil {
			caves[c2] = []string{c1}
		} else {
			caves[c2] = append(caves[c2], c1)
		}
	}
}
