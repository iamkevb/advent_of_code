package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type Point struct {
	x, y int
}
type Square struct {
	p    Point
	risk int
}

type PriorityQueue []Square

func (pq *PriorityQueue) insert(p Square) {
	*pq = append(*pq, p)
	sort.Sort(pq)
}

func (pq *PriorityQueue) dequeue() Square {
	old := *pq
	p := old[0]
	*pq = old[1:]
	return p
}

func (a PriorityQueue) Len() int           { return len(a) }
func (a PriorityQueue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PriorityQueue) Less(i, j int) bool { return a[i].risk < a[j].risk }

func main() {
	tile := readInput("./input.txt")
	graph := createGraph(tile)

	// Create a Distance Map Storing Distance from source Node to Each Nodes, Initialize all the valuse to Infinity
	riskMap := map[Point]int{}
	for y, row := range graph {
		for x, _ := range row {
			p := Point{x: x, y: y}
			riskMap[p] = math.MaxInt
		}
	}

	// For source Node mark distance as zero (weights should be non negative)
	source := Point{x: 0, y: 0}
	riskMap[source] = 0

	// Enqueue the source Node in min Priority queue
	priorityQueue := PriorityQueue{}
	priorityQueue.insert(Square{p: source, risk: graph[0][0]})

	visited := map[Point]bool{}

	x := len(graph) - 1
	y := len(graph[x]) - 1
	target := Point{x: x, y: y}

	for len(priorityQueue) > 0 {
		sq := priorityQueue.dequeue()

		if sq.p == target {
			break
		}
		visited[sq.p] = true

		adjacent := adjacentSquares(sq, graph)
		for _, adjacentSquare := range adjacent {
			if !visited[adjacentSquare.p] {
				pathRisk := riskMap[sq.p] + adjacentSquare.risk

				if riskMap[adjacentSquare.p] > pathRisk {
					riskMap[adjacentSquare.p] = pathRisk
					adjacentSquare.risk = pathRisk
					priorityQueue.insert(adjacentSquare)
				}
			}
		}
	}

	fmt.Println("total risk: ", riskMap[target])
}

func adjacentSquares(sq Square, graph [][]int) []Square {
	adj := []Square{}
	nx, ny := sq.p.x, sq.p.y-1
	if valid(nx, ny, graph) {
		north := Square{p: Point{x: nx, y: ny}, risk: graph[nx][ny]}
		adj = append(adj, north)
	}
	sx, sy := sq.p.x, sq.p.y+1
	if valid(sx, sy, graph) {
		south := Square{p: Point{x: sx, y: sy}, risk: graph[sx][sy]}
		adj = append(adj, south)
	}
	wx, wy := sq.p.x-1, sq.p.y
	if valid(wx, wy, graph) {
		west := Square{p: Point{x: wx, y: wy}, risk: graph[wx][wy]}
		adj = append(adj, west)
	}
	ex, ey := sq.p.x+1, sq.p.y
	if valid(ex, ey, graph) {
		east := Square{p: Point{x: ex, y: ey}, risk: graph[ex][ey]}
		adj = append(adj, east)
	}
	return adj
}

func valid(x, y int, graph [][]int) bool {
	return x >= 0 && x < len(graph) && y >= 0 && y < len(graph[x])
}

func createGraph(tile [][]int) [][]int {
	graph := [][]int{}

	tRows := len(tile)
	tCols := len(tile[0])

	for r := 0; r < tRows*5; r++ {
		graph = append(graph, make([]int, tCols*5))
	}

	for x := 0; x < tRows*5; x++ {
		for y := 0; y < tCols*5; y++ {
			xInc := x / tRows
			yInc := y / tCols
			tX := x % tRows
			tY := y % tCols
			val := tile[tX][tY] + xInc + yInc
			if val > 9 {
				val = val - 9
			}
			graph[x][y] = val
		}
	}
	return graph
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
