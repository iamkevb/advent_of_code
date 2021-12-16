package main

import (
	"bufio"
	"fmt"
	"os"
)

type Vertex struct {
	x, y int
}

func main() {
	graph := readInput("./tiny.txt")
	dists := computePaths(graph)

	bottomRight := Vertex{
		x: len(graph) - 1,
		y: len(graph[0]) - 1,
	}
	fmt.Println(dists[bottomRight])
}

func computePaths(graph [][]int) map[Vertex]int {
	dists := map[Vertex]int{
		{x: 0, y: 0}: 0,
	}
	findMinPath(0, 1, 0, graph, dists)
	findMinPath(1, 0, 0, graph, dists)
	return dists
}

func findMinPath(x, y, d int, graph [][]int, dists map[Vertex]int) {
	if !validCoord(x, y, graph) {
		return
	}
	vert := Vertex{x: x, y: y}
	dist := d + graph[x][y]
	if dists[vert] == 0 || dist < dists[vert] {
		dists[vert] = dist

		findMinPath(x, y-1, dist, graph, dists)
		findMinPath(x, y+1, dist, graph, dists)
		findMinPath(x+1, y, dist, graph, dists)
		findMinPath(x-1, y, dist, graph, dists)
	}
}

func validCoord(x, y int, graph [][]int) bool {
	if x == 0 && y == 0 {
		return false
	}

	return x >= 0 && y >= 0 && x < len(graph) && y < len(graph[x])
}

func readInput(path string) [][]int {
	var m [][]int = [][]int{}
	f, _ := os.Open(path)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for i, v := range line {
			row[i] = int(v - '0')
		}
		m = append(m, row)
	}
	return m
}
