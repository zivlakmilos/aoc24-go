package main

import (
	"fmt"
	"strings"
)

var nextLetterMap = map[byte]byte{
	'X': 'M',
	'M': 'A',
	'A': 'S',
}

var directions = []struct {
	x int
	y int
}{
	{
		x: -1,
		y: -1,
	},
	{
		x: 0,
		y: -1,
	},
	{
		x: 1,
		y: -1,
	},
	{
		x: -1,
		y: 0,
	},
	{
		x: 1,
		y: 0,
	},
	{
		x: -1,
		y: 1,
	},
	{
		x: 0,
		y: 1,
	},
	{
		x: 1,
		y: 1,
	},
}

func parseInput(input string) [][]byte {
	res := [][]byte{}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		l := make([]byte, len(line))
		for i := 0; i < len(line); i++ {
			l[i] = line[i]
		}
		res = append(res, l)
	}

	return res
}

func check(data [][]byte, ch byte, x, y, dirX, dirY int) bool {
	if y < 0 || y >= len(data) {
		return false
	}
	if x < 0 || x >= len(data[y]) {
		return false
	}

	if data[y][x] != ch {
		return false
	}

	if ch == 'S' {
		return true
	}

	x += dirX
	y += dirY

	if nextCh, ok := nextLetterMap[ch]; ok {
		return check(data, nextCh, x, y, dirX, dirY)
	}

	return false
}

func getCh(data [][]byte, x, y int) byte {
	if y < 0 || y >= len(data) {
		return 0
	}
	if x < 0 || x >= len(data[y]) {
		return 0
	}

	return data[y][x]
}

func checkX(data [][]byte, x, y int) bool {
	ch := getCh(data, x, y)

	if ch != 'A' {
		return false
	}

	ch1 := getCh(data, x-1, y-1)
	ch2 := getCh(data, x+1, y-1)
	ch3 := getCh(data, x-1, y+1)
	ch4 := getCh(data, x+1, y+1)
	if ((ch1 == 'M' && ch4 == 'S') || (ch1 == 'S' && ch4 == 'M')) &&
		((ch2 == 'M' && ch3 == 'S') || (ch2 == 'S' && ch3 == 'M')) {
		return true
	}

	return false
}

func solvePuzzle01() {
	input := getInput()

	res := 0

	data := parseInput(input)
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if data[i][j] == 'X' {
				for _, dir := range directions {
					if check(data, 'M', j+dir.x, i+dir.y, dir.x, dir.y) {
						res++
					}
				}
			}
		}
	}

	fmt.Printf("Number of XMAS: %d\n", res)
}

func solvePuzzle02() {
	input := getInput()

	res := 0

	data := parseInput(input)
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if checkX(data, j, i) {
				res++
			}
		}
	}

	fmt.Printf("Number of X-MAS: %d\n", res)
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
