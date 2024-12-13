package main

import (
	"fmt"
	"strings"
)

type Cell struct {
	value   byte
	visited bool
	id      int
}

type Area struct {
	area      int
	perimeter int
}

func parseInput(input string) [][]Cell {
	var res [][]Cell

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		row := make([]Cell, len(line))
		for i := 0; i < len(line); i++ {
			row[i] = Cell{
				value: line[i],
			}
		}
		res = append(res, row)
	}

	return res
}

func visit(r, c int, data [][]Cell, ch byte, id int) {
	if r < 0 || r >= len(data) || c < 0 || c >= len(data[r]) {
		return
	}
	if data[r][c].value != ch {
		return
	}
	if data[r][c].visited {
		return
	}

	data[r][c].id = id
	data[r][c].visited = true

	visit(r+1, c, data, ch, id)
	visit(r-1, c, data, ch, id)
	visit(r, c+1, data, ch, id)
	visit(r, c-1, data, ch, id)
}

func foundRegions(data [][]Cell) int {
	id := 0
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if data[i][j].visited {
				continue
			}

			visit(i, j, data, data[i][j].value, id)
			id++
		}
	}

	return id
}

func checkNeighbour(i, j int, data [][]Cell, id int) int {
	if i < 0 || i >= len(data) || j < 0 || j >= len(data[i]) {
		return 1
	}
	if data[i][j].id == id {
		return 0
	}

	return 1
}

func calculateAreas(data [][]Cell, count int) []Area {
	res := make([]Area, count)

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			res[data[i][j].id].area++

			res[data[i][j].id].perimeter += checkNeighbour(i+1, j, data, data[i][j].id)
			res[data[i][j].id].perimeter += checkNeighbour(i-1, j, data, data[i][j].id)
			res[data[i][j].id].perimeter += checkNeighbour(i, j+1, data, data[i][j].id)
			res[data[i][j].id].perimeter += checkNeighbour(i, j-1, data, data[i][j].id)
		}
	}

	return res
}

func solvePuzzle01() {
	input := getInput()
	data := parseInput(input)

	count := foundRegions(data)

	areas := calculateAreas(data, count)

	res := 0

	for _, area := range areas {
		res += area.area * area.perimeter
	}

	fmt.Printf("Total price of fencing: %d\n", res)
}

func main() {
	solvePuzzle01()
}
