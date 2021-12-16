package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"sort"
// )

// type Vertex struct {
// 	x, y int
// }

// func (v Vertex) valid(graph [][]int) bool {
// 	return v.x >= 0 && v.y >= 0 && v.x < len(graph) && v.y < len(graph[v.x])
// }

// type Square struct {
// 	v Vertex
// 	w int
// }

// func main() {
// 	tile := readInput("./input.txt")
// 	graph := createGraph(tile)
// 	dists := computePaths(graph)

// 	bottomRight := Vertex{
// 		x: len(graph) - 1,
// 		y: len(graph[0]) - 1,
// 	}
// 	fmt.Println(dists[bottomRight])
// }

// func createGraph(tile [][]int) [][]int {
// 	graph := [][]int{}

// 	tRows := len(tile)
// 	tCols := len(tile[0])

// 	for r := 0; r < tRows*5; r++ {
// 		graph = append(graph, make([]int, tCols*5))
// 	}

// 	for x := 0; x < tRows*5; x++ {
// 		for y := 0; y < tCols*5; y++ {
// 			xInc := x / tRows
// 			yInc := y / tCols
// 			tX := x % tRows
// 			tY := y % tCols
// 			val := tile[tX][tY] + xInc + yInc
// 			if val > 9 {
// 				val = val - 9
// 			}
// 			graph[x][y] = val
// 		}
// 	}
// 	return graph
// }

// func computePaths(graph [][]int) map[Vertex]int {
// 	source := Vertex{x: 0, y: 0}
// 	dists := map[Vertex]int{
// 		source: 0,
// 	}

// 	findMinPath(0, Vertex{x: 0, y: 1}, source, graph, dists)
// 	findMinPath(0, Vertex{x: 1, y: 0}, source, graph, dists)
// 	return dists
// }

// var count int = 0

// func findMinPath(d int, next, last Vertex, graph [][]int, dists map[Vertex]int) {
// 	if !next.valid(graph) {
// 		return
// 	}

// 	count++
// 	if count%10000 == 0 {
// 		fmt.Printf("visited %d squares. Distances computed: %d\n", count, len(dists))
// 	}

// 	dist := d + graph[next.x][next.y]

// 	if dists[next] == 0 || dist < dists[next] {
// 		dists[next] = dist

// 		v := nextSquares(next, graph)
// 		for _, val := range v {
// 			findMinPath(dist, val, next, graph, dists)
// 		}
// 	}
// }

// func nextSquares(v Vertex, graph [][]int) []Vertex {
// 	squares := SquareSlice{}

// 	vert := Vertex{x: v.x, y: v.y - 1}
// 	if vert.valid(graph) {
// 		squares = append(squares, Square{v: vert, w: graph[vert.x][vert.y]})
// 	}

// 	vert = Vertex{x: v.x, y: v.y + 1}
// 	if vert.valid(graph) {
// 		squares = append(squares, Square{v: vert, w: graph[vert.x][vert.y]})
// 	}

// 	vert = Vertex{x: v.x - 1, y: v.y}
// 	if vert.valid(graph) {
// 		squares = append(squares, Square{v: vert, w: graph[vert.x][vert.y]})
// 	}

// 	vert = Vertex{x: v.x + 1, y: v.y}
// 	if vert.valid(graph) {
// 		squares = append(squares, Square{v: vert, w: graph[vert.x][vert.y]})
// 	}

// 	sort.Sort(squares)
// 	vertices := []Vertex{}
// 	for _, s := range squares {
// 		vertices = append(vertices, s.v)
// 	}
// 	return vertices
// }

// type SquareSlice []Square

// func (a SquareSlice) Len() int           { return len(a) }
// func (a SquareSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a SquareSlice) Less(i, j int) bool { return a[i].w > a[j].w }

// func readInput(path string) [][]int {
// 	var m [][]int = [][]int{}
// 	f, _ := os.Open(path)
// 	scanner := bufio.NewScanner(f)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		row := make([]int, len(line))
// 		for i, v := range line {
// 			row[i] = int(v - '0')
// 		}
// 		m = append(m, row)
// 	}
// 	return m
// }
