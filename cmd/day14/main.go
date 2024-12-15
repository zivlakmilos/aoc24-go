package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Robot struct {
	x  int
	y  int
	vx int
	vy int
}

func eucledianReminder(a, b int) int {
	r := a % b

	if r < 0 {
		return r + b
	}

	return r
}

func parseRobot(str string) Robot {
	var robot Robot

	data := strings.Split(str, " ")
	p := strings.Split(data[0], ",")
	v := strings.Split(data[1], ",")

	robot.x, _ = strconv.Atoi(strings.Split(p[0], "=")[1])
	robot.y, _ = strconv.Atoi(p[1])

	robot.vx, _ = strconv.Atoi(strings.Split(v[0], "=")[1])
	robot.vy, _ = strconv.Atoi(v[1])

	return robot
}

func parseInput(input string) []Robot {
	var res []Robot

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		res = append(res, parseRobot(line))
	}

	return res
}

func simulate(robot *Robot, w, h, t int) {
	robot.x += robot.vx * t
	robot.y += robot.vy * t

	robot.x = eucledianReminder(robot.x, w)
	robot.y = eucledianReminder(robot.y, h)
}

func solvePuzzle01() {
	width := 101
	height := 103

	input := getInput()
	robots := parseInput(input)

	for idx := range robots {
		simulate(&robots[idx], width, height, 100)
	}

	q1 := 0
	q2 := 0
	q3 := 0
	q4 := 0

	for i := 0; i < height; i++ {
		if i == height/2 {
			continue
		}

		for j := 0; j < width; j++ {
			if j == width/2 {
				continue
			}

			for _, r := range robots {
				if r.x == j && r.y == i {
					if r.y < height/2 {
						if r.x < width/2 {
							q1++
						} else if r.x > width/2 {
							q2++
						}
					} else if r.y > height/2 {
						if r.x < width/2 {
							q3++
						} else if r.x > width/2 {
							q4++
						}
					}
				}
			}
		}
	}

	res := q1 * q2 * q3 * q4
	fmt.Printf("Result: %d\n", res)
}

func main() {
	solvePuzzle01()
}
