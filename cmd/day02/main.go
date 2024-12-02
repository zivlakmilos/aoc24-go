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

func processLine(line string) bool {
	values := parseLine(line)

	prevDiff := 0
	for i := 1; i < len(values); i++ {
		diff := values[i-1] - values[i]
		if (diff < 0 && prevDiff > 0) || (diff > 0 && prevDiff < 0) {
			return false
		}

		if diff > 3 || diff < -3 || diff == 0 {
			return false
		}

		prevDiff = diff
	}

	return true
}

func solvePuzzle01() {
	input := getInput()

	lines := strings.Split(input, "\n")

	res := 0
	for _, line := range lines {
		if processLine(line) {
			res++
		}
	}

	fmt.Printf("Safe repports: %d\n", res)
}

func main() {
	solvePuzzle01()
}
