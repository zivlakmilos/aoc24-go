package main

import (
	"fmt"
	"strings"
)

type Coord struct {
	x int
	y int
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}

	return num
}

func isValidCoord(x, y, w, h int) bool {
	if x < 0 || x >= w || y < 0 || y >= h {
		return false
	}

	return true
}

func parseInput(input string) (map[byte][]Coord, int, int) {
	res := map[byte][]Coord{}

	lines := strings.Split(input, "\n")
	w := len(lines[0])
	h := len(lines)

	for y := range lines {
		for x := range lines[y] {
			if lines[y][x] != '.' {
				ch := lines[y][x]
				res[ch] = append(res[ch], Coord{
					x: x,
					y: y,
				})
			}
		}
	}

	return res, w, h
}

func calcAntinodes(a, b Coord, w, h int) []Coord {
	var res []Coord

	diffX := a.x - b.x
	diffY := a.y - b.y

	x1 := a.x + diffX
	y1 := a.y + diffY
	x2 := b.x - diffX
	y2 := b.y - diffY

	if isValidCoord(x1, y1, w, h) {
		res = append(res, Coord{
			x: x1,
			y: y1,
		})
	}

	if isValidCoord(x2, y2, w, h) {
		res = append(res, Coord{
			x: x2,
			y: y2,
		})
	}

	return res
}

func solvePuzzle01() {
	input := getInput()
	data, w, h := parseInput(input)

	set := map[string]struct{}{}
	for _, val := range data {
		for i := 0; i < len(val); i++ {
			for j := i + 1; j < len(val); j++ {
				antinodes := calcAntinodes(val[i], val[j], w, h)
				for _, antinode := range antinodes {
					key := fmt.Sprintf("%d-%d", antinode.x, antinode.y)
					set[key] = struct{}{}
				}
			}
		}
	}

	res := len(set)
	fmt.Printf("Number of antinodes: %d\n", res)
}

func main() {
	solvePuzzle01()
}
