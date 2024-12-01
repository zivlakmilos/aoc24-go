package main

import (
	"fmt"
	"slices"
	"strings"
)

func parseLine(line string) (int, int) {
	var l int
	var r int

	fmt.Sscanf(line, "%d %d", &l, &r)

	return l, r
}

func solvePuzzle01() {
	input := getInput()

	var left []int
	var right []int

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		l, r := parseLine(line)

		left = append(left, l)
		right = append(right, r)
	}

	slices.Sort(left)
	slices.Sort(right)

	res := 0

	for i := range left {
		l := left[i]
		r := right[i]
		diff := 0

		if l > r {
			diff = l - r
		} else {
			diff = r - l
		}

		res += diff
	}

	fmt.Printf("Total distance between lists: %d\n", res)
}

func solvePuzzle02() {
	input := getInput()

	var left []int
	rightCount := map[int]int{}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		l, r := parseLine(line)

		left = append(left, l)
		rightCount[r]++
	}

	res := 0

	for i := range left {
		l := left[i]
		r := rightCount[l]

		res += l * r
	}

	fmt.Printf("Similarity score: %d\n", res)
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
