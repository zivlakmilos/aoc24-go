package main

import (
	"fmt"
	"strconv"
	"strings"
)

const Inf = ^uint(0x00)

type Coord struct {
	x int
	y int
}

func parseCoord(line string, sep string) Coord {
	var res Coord

	str := strings.Split(line, ",")
	for idx, s := range str {
		str2 := strings.Split(s, sep)
		val, _ := strconv.Atoi(str2[1])
		if idx == 0 {
			res.x = val
		} else {
			res.y = val
		}
	}

	return res
}

func parseInput(input string) (Coord, Coord, Coord) {
	lines := strings.Split(input, "\n")

	a := parseCoord(lines[0], "+")
	b := parseCoord(lines[1], "+")
	prize := parseCoord(lines[2], "=")

	return a, b, prize
}

func solveOne01(input string) uint {
	a, b, prize := parseInput(input)

	res := Inf

	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100; j++ {
			x := a.x*i + b.x*j
			y := a.y*i + b.y*j
			if x == prize.x && y == prize.y {
				tokens := uint(i*3 + j)
				if tokens < res {
					res = tokens
				}
			}
		}
	}

	if res == Inf {
		return 0
	}

	return res
}

func solvePuzzle01() {
	input := getInput()

	res := uint(0)

	data := strings.Split(input, "\n\n")
	for _, d := range data {
		res += solveOne01(d)
	}

	fmt.Printf("Result: %d\n", res)
}

func main() {
	solvePuzzle01()
}
