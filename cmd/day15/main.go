package main

import (
	"fmt"
	"strings"
)

func parseInput(input string) ([][]byte, []byte) {
	var pos [][]byte
	var moves []byte

	data := strings.Split(input, "\n\n")

	lines := strings.Split(data[0], "\n")
	for _, line := range lines {
		pos = append(pos, []byte(line))
	}

	lines = strings.Split(data[1], "\n")
	for _, line := range lines {
		moves = append(moves, []byte(line)...)
	}

	return pos, moves
}

func enlargeMap(pos [][]byte) [][]byte {
	var res [][]byte

	for i := 0; i < len(pos); i++ {
		var row []byte
		for j := 0; j < len(pos[i]); j++ {
			ch1 := pos[i][j]
			ch2 := pos[i][j]
			switch ch1 {
			case '@':
				ch2 = '.'
			case 'O':
				ch1 = '['
				ch2 = ']'
			}
			row = append(row, ch1)
			row = append(row, ch2)
		}
		res = append(res, row)
	}

	return res
}

func findRobot(pos [][]byte) (int, int) {
	for i := range pos {
		for j := range pos[i] {
			if pos[i][j] == '@' {
				return j, i
			}
		}
	}

	return 0, 0
}

func calcRobotPos(x, y int, move byte) (int, int) {
	switch move {
	case '^':
		y--
	case '>':
		x++
	case 'v':
		y++
	case '<':
		x--
	}

	return x, y
}

func makeMove(pos [][]byte, x, y int, move byte, ch byte) bool {
	if pos[y][x] == '#' {
		return false
	}
	if pos[y][x] == '.' {
		return true
	}

	nx := x
	ny := y

	switch move {
	case '^':
		ny--
	case '>':
		nx++
	case 'v':
		ny++
	case '<':
		nx--
	}

	if pos[y][x] == '[' && (move == '^' || move == 'v') {
		if makeMove(pos, nx, ny, move, pos[y][x]) && makeMove(pos, nx+1, ny, move, pos[y][x+1]) {
			pos[ny][nx] = pos[y][x]
			pos[y][x] = '.'
			pos[ny][nx+1] = pos[y][x+1]
			pos[y][x+1] = '.'
			return true
		}
	} else if pos[y][x] == ']' && (move == '^' || move == 'v') {
		if makeMove(pos, nx, ny, move, pos[y][x]) && makeMove(pos, nx-1, ny, move, pos[y][x-1]) {
			pos[ny][nx] = pos[y][x]
			pos[y][x] = '.'
			pos[ny][nx-1] = pos[y][x-1]
			pos[y][x-1] = '.'
			return true
		}
	} else {
		if makeMove(pos, nx, ny, move, pos[y][x]) {
			pos[ny][nx] = pos[y][x]
			pos[y][x] = '.'
			return true
		}
	}

	return false
}

func isMovePossible(pos [][]byte, x, y int, move byte, ch byte) bool {
	if pos[y][x] == '#' {
		return false
	}
	if pos[y][x] == '.' {
		return true
	}

	nx := x
	ny := y

	switch move {
	case '^':
		ny--
	case '>':
		nx++
	case 'v':
		ny++
	case '<':
		nx--
	}

	if pos[y][x] == '[' && (move == '^' || move == 'v') {
		if isMovePossible(pos, nx, ny, move, pos[y][x]) && isMovePossible(pos, nx+1, ny, move, pos[y][x+1]) {
			return true
		}
	} else if pos[y][x] == ']' && (move == '^' || move == 'v') {
		if isMovePossible(pos, nx, ny, move, pos[y][x]) && isMovePossible(pos, nx-1, ny, move, pos[y][x-1]) {
			return true
		}
	} else {
		if isMovePossible(pos, nx, ny, move, pos[y][x]) {
			return true
		}
	}

	return false
}

func printMap(pos [][]byte) {
	for i := range pos {
		for j := range pos[i] {
			fmt.Printf("%c", pos[i][j])
		}
		fmt.Printf("\n")
	}
}

func calcGps(pos [][]byte) int {
	res := 0

	for i := range pos {
		for j := range pos[i] {
			if pos[i][j] == 'O' || pos[i][j] == '[' {
				res += 100*i + j
			}
		}
	}

	return res
}

func solvePuzzle01() {
	input := getInput()
	pos, moves := parseInput(input)
	rx, ry := findRobot(pos)

	for _, move := range moves {
		if makeMove(pos, rx, ry, move, '@') {
			rx, ry = calcRobotPos(rx, ry, move)
		}
	}

	res := calcGps(pos)
	fmt.Printf("Result: %d\n", res)
}

func solvePuzzle02() {
	input := getInput()
	pos, moves := parseInput(input)
	pos = enlargeMap(pos)
	rx, ry := findRobot(pos)

	for _, move := range moves {
		if isMovePossible(pos, rx, ry, move, '@') {
			makeMove(pos, rx, ry, move, '@')
			rx, ry = calcRobotPos(rx, ry, move)
		}
	}

	printMap(pos)

	res := calcGps(pos)
	fmt.Printf("Result: %d\n", res)
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
