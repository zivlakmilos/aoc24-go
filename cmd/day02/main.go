package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseLine(line string) []int {
	values := strings.Split(line, " ")

	res := make([]int, len(values))
	for i := range values {
		res[i], _ = strconv.Atoi(values[i])
	}

	return res
}

func checkDiff(v1, v2, prevDiff int) bool {
	diff := v1 - v2

	if (diff < 0 && prevDiff > 0) || (diff > 0 && prevDiff < 0) {
		return false
	}

	if diff > 3 || diff < -3 || diff == 0 {
		return false
	}

	return true
}

func processValues(values []int) bool {
	prevDiff := 0
	for i := 1; i < len(values); i++ {
		if !checkDiff(values[i], values[i-1], prevDiff) {
			return false
		}

		prevDiff = values[i] - values[i-1]
	}

	return true
}

func processLine(line string, tolerance int) bool {
	values := parseLine(line)

	if processValues(values) {
		return true
	}

	if tolerance > 0 {
		for i := 0; i < len(values); i++ {
			var vals []int
			vals = append(vals, values[:i]...)
			vals = append(vals, values[i+1:]...)
			if processValues(vals) {
				return true
			}
		}
	}

	return false

	prevDiff := 0
	i := 0
	j := 1
	bad := 0
	for j < len(values) {
		if bad > tolerance {
			return false
		}

		diff := values[i] - values[j]
		if !checkDiff(values[i], values[j], prevDiff) {
			bad++
			if j+1 < len(values) && checkDiff(values[i], values[j+1], prevDiff) {
				j++
			} else if i == 0 {
				i = 1
				j = 2
				continue
			} else if j == len(values)-1 {
			} else {
				fmt.Printf("%#v, (%d, %d)\n", values, i, j)
				return false
			}

			diff = values[i] - values[j]

			i = j
			j++
			continue
		}

		i = j
		j++
		prevDiff = diff
	}

	return bad <= tolerance
}

func solvePuzzle01() {
	input := getInput()

	lines := strings.Split(input, "\n")

	res := 0
	for _, line := range lines {
		if processLine(line, 0) {
			res++
		}
	}

	fmt.Printf("Safe repports: %d\n", res)
}

func solvePuzzle02() {
	input := getInput()

	lines := strings.Split(input, "\n")

	res := 0
	for _, line := range lines {
		if processLine(line, 1) {
			res++
		}
	}

	fmt.Printf("Safe repports: %d\n", res)
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
