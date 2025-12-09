package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Box struct {
	x, y, z int
}

type Edge struct {
	b1, b2 int
	dst    int64
}

type UF struct {
	parent, size []int
}

func (uf *UF) Find(a int) int {
	for uf.parent[a] != a {
		uf.parent[a] = uf.parent[uf.parent[a]]
		a = uf.parent[a]
	}
	return a
}

func (uf *UF) Union(a, b int) bool {
	ra := uf.Find(a)
	rb := uf.Find(b)
	if ra == rb {
		return false
	}
	uf.parent[rb] = ra
	uf.size[ra] += uf.size[rb]
	return true
}

const TEST = "./input-test.txt"
const DATA = "./input.txt"

func main() {
	input := DATA

	switch input {
	case TEST:
		boxes := readInput(TEST)
		part1(boxes, 10)
		part2(boxes)

	case DATA:
		boxes := readInput(DATA)
		part1(boxes, 1000)
		part2(boxes)
	}
}

func makeEdges(boxes []Box) []Edge {
	edges := []Edge{}
	for i, ib := range boxes {
		for j := i + 1; j < len(boxes); j++ {
			jb := boxes[j]
			dx := int64(ib.x - jb.x)
			dy := int64(ib.y - jb.y)
			dz := int64(ib.z - jb.z)
			dst := dx*dx + dy*dy + dz*dz
			edges = append(edges, Edge{b1: i, b2: j, dst: dst})
		}
	}
	return edges
}

func newUnionFind(n int) *UF {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &UF{parent, size}
}

func part1(boxes []Box, connections int) {
	edges := makeEdges(boxes)
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dst < edges[j].dst
	})

	uf := newUnionFind(len(boxes))

	for i := 0; i < connections; i++ {
		edge := edges[i]
		uf.Union(edge.b1, edge.b2)
	}

	//find number of nodes that map back to each root.
	counts := map[int]int{}
	for i := 0; i < len(boxes); i++ {
		root := uf.Find(i)
		counts[root]++
	}
	sizes := []int{}
	for _, v := range counts {
		sizes = append(sizes, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	product := sizes[0] * sizes[1] * sizes[2]
	fmt.Println("Part 1:", product)
}

func part2(boxes []Box) {
	edges := makeEdges(boxes)

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dst < edges[j].dst
	})

	uf := newUnionFind(len(boxes))

	lonely := len(boxes)
	lastA, lastB := 0, 0

	for _, e := range edges {
		if uf.Union(e.b1, e.b2) {
			lonely--
			lastA, lastB = e.b1, e.b2
			if lonely == 1 {
				break
			}

			// THIS ALSO WORKS. Ensure all edges are connected to the same number of nodes (all)
			// counts := map[int]int{}
			// for i := 0; i < len(boxes); i++ {
			// 	root := uf.Find(i)
			// 	counts[root]++
			// }
			// if len(counts) == 1 {
			// 	break
			// }
		}
	}
	fmt.Println("Part 2:", boxes[lastA].x*boxes[lastB].x)

}

func readInput(path string) []Box {
	f, _ := os.ReadFile(path)
	boxes := []Box{}
	scanner := bufio.NewScanner(bytes.NewReader(f))
	for scanner.Scan() {
		line := scanner.Text()
		coord := strings.Split(line, ",")
		x, _ := strconv.Atoi(coord[0])
		y, _ := strconv.Atoi(coord[1])
		z, _ := strconv.Atoi(coord[2])
		boxes = append(boxes, Box{x, y, z})

	}
	return boxes
}
