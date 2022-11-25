package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	value, depth int
}

func main() {
	// path := "./test.txt"
	path := "./input.txt"
	numbers := readInput(path)

	max := 0

	for i := 0; i < len(numbers)-1; i++ {
		n1 := numbers[i]
		for j := i; j < len(numbers); j++ {
			n2 := numbers[j]

			n1c := make([]Node, len(n1))
			n2c := make([]Node, len(n2))
			copy(n1c, n1)
			copy(n2c, n2)

			sum1 := add(n1c, n2c)
			copy(n1c, n1)
			copy(n2c, n2)
			sum2 := add(n2c, n1c)

			mag1 := magnitude(sum1)
			mag2 := magnitude(sum2)
			if mag1 > max {
				max = mag1
			}
			if mag2 > max {
				max = mag2
			}
		}

	}

	fmt.Println(max)
}

func magnitude(n []Node) int {
	for depth := 4; depth > 0; {
		for i := 0; i < len(n); i++ {
			if n[i].depth == depth {
				l := n[i].value
				r := n[i+1].value
				n[i].value = l*3 + r*2
				n[i].depth--
				n = append(n[:i+1], n[i+2:]...)
			}
		}
		depth--
	}
	//only one node left?
	return n[0].value
}

func add(n1, n2 []Node) []Node {
	nodes := append(n1, n2...)
	//increase depth of all by one.
	for i := range nodes {
		nodes[i].depth += 1
	}
	return reduce(nodes)
}

func reduce(n []Node) []Node {
	finished := false
	for !finished {
		n, finished = explode(n)
		if !finished {
			continue
		}

		n, finished = split(n)
	}
	return n
}

func split(n []Node) ([]Node, bool) {
	finished := true
	for i := range n {
		if n[i].value >= 10 {
			val := n[i].value
			n[i].value = val / 2
			n[i].depth++
			n = append(n, Node{}) //ensure room for one more
			copy(n[i+2:], n[i+1:])
			n[i+1] = Node{value: val - n[i].value, depth: n[i].depth}
			finished = false
			break
		}
	}
	return n, finished
}

func explode(n []Node) ([]Node, bool) {
	finished := true
	for i := range n {
		if n[i].depth >= 5 {
			if i > 0 {
				n[i-1].value += n[i].value
			}
			if i+2 < len(n) {
				n[i+2].value += n[i+1].value
			}

			n[i].value = 0
			n[i].depth -= 1
			n = append(n[:i+1], n[i+2:]...)
			finished = false
			break
		}
	}
	return n, finished
}

func parseNode(text string) []Node {
	nodes := []Node{}

	depth := 0

	for _, v := range text {
		switch v {
		case '[':
			depth++
		case ']':
			depth--
		case ',':
			continue
		default:
			var val int = int(v - '0')
			n := Node{
				value: val,
				depth: depth,
			}
			nodes = append(nodes, n)
		}
	}
	return nodes
}

func readInput(path string) [][]Node {
	var numbers [][]Node = [][]Node{}
	f, _ := os.Open(path)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		nodes := parseNode(line)
		numbers = append(numbers, nodes)
	}
	return numbers
}
