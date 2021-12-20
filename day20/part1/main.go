package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Pixel struct {
	x, y int
}

var algo []rune

var bgLit = false

func main() {
	path := "./test.txt"
	path = "./input.txt"
	img := readInput(path)

	for i := 0; i < 50; i++ {
		img = enhance(img)
	}

	lighted := 0
	for _, px := range img {
		if px {
			lighted += 1
		}
	}
	fmt.Printf("Lighted pixels: %d\n", lighted)
}

func enhance(img map[Pixel]bool) map[Pixel]bool {
	min, max := minMaxXY(img)
	enhanced := make(map[Pixel]bool)

	for x := min.x - 1; x <= max.x+1; x++ {
		for y := min.y - 1; y <= max.y+1; y++ {
			enhanced[Pixel{x, y}] = enhancePixel(x, y, img)
		}
	}
	if bgLit {
		bgLit = algo[len(algo)-1] == '#'
	} else {
		bgLit = algo[0] == '#'
	}
	return enhanced
}

func enhancePixel(x, y int, img map[Pixel]bool) bool {
	idx := 0
	for xc := x - 1; xc <= x+1; xc++ {
		for yc := y - 1; yc <= y+1; yc++ {
			p, found := img[Pixel{xc, yc}]
			if !found {
				p = bgLit
			}

			idx = idx << 1
			if p {
				idx = idx + 1
			}
		}
	}

	return []rune(algo)[idx] == '#'
}

func minMaxXY(img map[Pixel]bool) (Pixel, Pixel) {
	min, max := Pixel{math.MaxInt, math.MaxInt}, Pixel{math.MinInt, math.MinInt}
	for v := range img {
		if v.x < min.x {
			min.x = v.x
		}
		if v.x > max.x {
			max.x = v.x
		}
		if v.y < min.y {
			min.y = v.y
		}
		if v.y > max.y {
			max.y = v.y
		}
	}
	return min, max
}

func writeImg(img map[Pixel]bool) {
	min, max := minMaxXY(img)
	for x := min.x; x <= max.x; x++ {
		for y := min.y; y <= max.y; y++ {
			if img[Pixel{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func readInput(path string) map[Pixel]bool {
	f, _ := os.Open(path)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	algo = []rune(scanner.Text())
	return readImage(scanner)
}

func readImage(scanner *bufio.Scanner) map[Pixel]bool {
	img := map[Pixel]bool{}
	x := 0
	for scanner.Scan() {
		t := scanner.Text()
		if len(t) == 0 {
			continue
		}
		for y, v := range t {
			px := Pixel{x, y}
			img[px] = v == '#'
		}
		x++
	}
	return img
}
