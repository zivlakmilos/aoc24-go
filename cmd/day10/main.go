package main

import (
	"fmt"
	"strings"
)

var neighbours = []struct {
	i int
	j int
}{
	{
		i: -1,
		j: 0,
	},
	{
		i: 1,
		j: 0,
	},
	{
		i: 0,
		j: -1,
	},
	{
		i: 0,
		j: 1,
	},
}

func parseInput(input string) [][]byte {
	var res [][]byte

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		var row []byte
		for i := 0; i < len(line); i++ {
			row = append(row, line[i]-'0')
		}
		res = append(res, row)
	}

	return res
}

func getCell(data [][]byte, i, j int) int {
	if i < 0 || i >= len(data) {
		return -1
	}
	if j < 0 || j >= len(data[i]) {
		return -1
	}

	return int(data[i][j])
}

func countTrails(data [][]byte, i, j int, curr byte, set map[string]struct{}) int {
	if i < 0 || i >= len(data) || j < 0 || j >= len(data[i]) {
		return 0
	}
	if data[i][j] != curr {
		return 0
	}

	key := fmt.Sprintf("%d-%d", i, j)
	if _, ok := set[key]; ok {
		return 0
	}
	set[key] = struct{}{}

	if data[i][j] == 9 {
		return 1
	}

	res := 0

	for _, n := range neighbours {
		res += countTrails(data, i+n.i, j+n.j, curr+1, set)
	}

	return res
}

func solvePuzzle01() {
	input := getInput()
	data := parseInput(input)

	res := 0

	for i := range data {
		for j := range data[i] {
			set := map[string]struct{}{}
			res += countTrails(data, i, j, 0, set)
		}
	}

	fmt.Printf("Result: %d\n", res)
}

func main() {
	solvePuzzle01()
}
