package main

import (
	"fmt"
	"strings"
)

type Cell struct {
	visited bool
	value   byte
}

func parseInput(input string) ([][]Cell, int, int) {
	res := [][]Cell{}
	x := 0
	y := 0

	lines := strings.Split(input, "\n")
	for r, line := range lines {
		row := make([]Cell, len(line))

		for i := range line {
			if line[i] == '^' {
				x = i
				y = r
			}
			row[i].value = line[i]
			row[i].visited = false
		}

		res = append(res, row)
	}

	return res, x, y
}

func getNextCoords(x, y int, dir byte) (int, int) {
	switch dir {
	case '^':
		y--
	case '>':
		x++
	case 'V':
		y++
	case '<':
		x--
	}

	return x, y
}

func rotate(dir byte) byte {
	switch dir {
	case '^':
		return '>'
	case '>':
		return 'V'
	case 'V':
		return '<'
	case '<':
		return '^'
	}

	return dir
}

func isValidPos(matrix [][]Cell, x, y int) bool {
	if x < 0 || y < 0 || x >= len(matrix) || y >= len(matrix[1]) {
		return false
	}

	return true
}

func makeMove(matrix [][]Cell, x, y int, dir byte) (int, int, bool) {
	nx, ny := getNextCoords(x, y, dir)
	if !isValidPos(matrix, nx, ny) {
		return nx, ny, true
	}

	if matrix[ny][nx].value == '#' {
		return x, y, false
	}

	return nx, ny, true
}

func countVisited(matrix [][]Cell) int {
	res := 0

	for _, row := range matrix {
		for _, cell := range row {
			if cell.visited {
				res++
			}
		}
	}

	return res
}

func solvePuzzle01() {
	input := getInput()
	matrix, x, y := parseInput(input)
	dir := byte('^')

	for isValidPos(matrix, x, y) {
		matrix[y][x].visited = true
		var ok bool
		x, y, ok = makeMove(matrix, x, y, dir)
		if !ok {
			dir = rotate(dir)
		}
	}

	res := countVisited(matrix)
	fmt.Printf("Result: %d\n", res)
}

func main() {
	solvePuzzle01()
}
