package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Game struct {
	Id    int
	Green int
	Red   int
	Blue  int
}

func main() {
	lines := readInput("./day2.input")

	// part1(lines)
	part2(lines)
}

func part2(lines []string) {
	games := []Game{}
	for _, l := range lines {
		g := parseGame(l, maxValue)
		games = append(games, g)
	}

	total := 0
	for _, g := range games {
		fmt.Printf("%+v\n", g)
		total += g.Red * g.Green * g.Blue
	}
	fmt.Println("Part 2 Answer: ", total)
}

func part1(lines []string) {
	games := []Game{}
	for _, l := range lines {
		g := parseGame(l, maxValue)
		games = append(games, g)
	}

	red := 12
	green := 13
	blue := 14
	total := 0
	for _, g := range games {
		if g.Red <= red && g.Green <= green && g.Blue <= blue {
			total += g.Id
		}
	}
	fmt.Println("Answer: ", total)
}

var reGame = regexp.MustCompile(`Game (\d+)`)
var reRed = regexp.MustCompile(`(\d+) red`)
var reGreen = regexp.MustCompile(`(\d+) green`)
var reBlue = regexp.MustCompile(`(\d+) blue`)

func parseGame(line string, fn func(vals [][]string) int) Game {
	id := reGame.FindAllStringSubmatch(line, -1)
	red := reRed.FindAllStringSubmatch(line, -1)
	green := reGreen.FindAllStringSubmatch(line, -1)
	blue := reBlue.FindAllStringSubmatch(line, -1)

	return Game{
		Id:    fn(id),
		Red:   fn(red),
		Green: fn(green),
		Blue:  fn(blue),
	}
}

func maxValue(vals [][]string) int {
	mx := 0
	for _, v := range vals {
		i, _ := strconv.Atoi(v[1])
		mx = max(mx, i)
	}
	return mx
}

func readInput(path string) []string {
	f, _ := os.Open(path)
	scanner := bufio.NewScanner(f)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}
