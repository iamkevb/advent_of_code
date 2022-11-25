package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Dot struct {
	x, y int
}
type Fold struct {
	axis  string
	value int
}

type Paper map[Dot]bool

var paper Paper = Paper{}
var folds []Fold

func main() {
	readInput("./input.txt")
	for _, fold := range folds {
		paper = foldPaper(fold)
		fmt.Println(len(paper))
	}
}

func foldPaper(fold Fold) Paper {
	if fold.axis == "x" {
		return foldHorizontally(fold)
	} else {
		return foldVertically(fold)
	}
}

func foldHorizontally(fold Fold) Paper {
	var p Paper = Paper{}
	for dot, _ := range paper {
		x, y := dot.x, dot.y
		if dot.x > fold.value {
			x = fold.value - (dot.x - fold.value)
		}
		p[Dot{x: x, y: y}] = true
	}
	return p
}

func foldVertically(fold Fold) Paper {
	var p Paper = Paper{}
	for dot, _ := range paper {
		x, y := dot.x, dot.y
		if dot.y > fold.value {
			y = fold.value - (dot.y - fold.value)
		}
		p[Dot{x: x, y: y}] = true
	}
	return p
}

func readInput(path string) {
	f, _ := os.Open(path)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, "fold") {
			var axis string
			var value int
			fmt.Sscanf(line, "fold along %1s=%d", &axis, &value)
			folds = append(folds, Fold{axis: axis, value: value})
		} else {
			var x, y int
			fmt.Sscanf(line, "%d,%d", &x, &y)
			paper[Dot{x: x, y: y}] = true
		}
	}
}
