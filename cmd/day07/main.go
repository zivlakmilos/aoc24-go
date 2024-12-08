package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseLine(line string) (int, []int) {
	var target int
	var values []int

	strings := strings.Split(line, " ")
	for idx, str := range strings {
		if idx == 0 {
			target, _ = strconv.Atoi(str[:len(str)-1])
			continue
		}

		val, _ := strconv.Atoi(str)
		values = append(values, val)
	}

	return target, values
}

func checkIsValid(target int, values []int) bool {
	if len(values) == 0 {
		return target == 0
	}

	if target <= 0 {
		return false
	}

	val := values[len(values)-1]
	if checkIsValid(target-val, values[:len(values)-1]) {
		return true
	}
	if target%val == 0 && checkIsValid(target/val, values[:len(values)-1]) {
		return true
	}

	return false
}

func solvePuzzle01() {
	input := getInput()

	res := 0

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		target, values := parseLine(line)
		if checkIsValid(target, values) {
			res += target
		}
	}

	fmt.Printf("Total calibration results: %d\n", res)
}

func main() {
	solvePuzzle01()
}
