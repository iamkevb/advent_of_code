package main

import (
	"bufio"
	"errors"
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
	path := []string{}
	paths, _ := findPaths(path, "start", "end")
	return paths
}

func findPaths(currentPath []string, start, end string) ([][]string, error) {
	if !canVisit(currentPath, start) {
		return [][]string{}, errors.New("not a valid path")
	}

	path := append(currentPath, start)
	if start == end {
		return [][]string{path}, nil
	}

	node := caves[start]
	completePaths := [][]string{}

	for _, n := range node {
		p := make([]string, len(path))
		copy(p, path)

		completed, err := findPaths(p, n, end)
		if err == nil {
			completePaths = append(completePaths, completed...)
		}

	}

	return completePaths, nil
}

func canVisit(arr []string, value string) bool {
	if contains(arr, "end") {
		return false
	}
	if value == "start" {
		return !contains(arr, value)
	}
	if strings.ToUpper(value) == value {
		return true
	}

	visits := map[string]int{}

	visits[value] = visits[value] + 1
	for _, s := range arr {
		if strings.ToUpper(s) == s {
			continue
		}
		visits[s] = visits[s] + 1
	}

	if visits[value] == 3 {
		return false
	}

	doubleVisit := false
	for _, v := range visits {
		if doubleVisit && v >= 2 {
			return false
		}
		if v == 2 {
			doubleVisit = true
		}

	}
	return true
}

func contains(arr []string, val string) bool {
	for _, v := range arr {
		if v == val {
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
