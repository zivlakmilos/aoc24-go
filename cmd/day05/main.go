package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseRule(rules map[int][]int, line string) {
	parsed := strings.Split(line, "|")

	key, _ := strconv.Atoi(parsed[0])
	val, _ := strconv.Atoi(parsed[1])

	rules[key] = append(rules[key], val)
}

func checkLine(rules map[int][]int, line string) int {
	res := 0

	set := map[int]struct{}{}

	data := strings.Split(line, ",")
	for i := range data {
		val, _ := strconv.Atoi(data[i])

		for _, val2 := range rules[val] {
			if _, ok := set[val2]; ok {
				return 0
			}
		}

		if i == len(data)/2 {
			res = val
		}

		set[val] = struct{}{}
	}

	return res
}

func solvePuzzle01() {
	input := getInput()

	res := 0

	rules := map[int][]int{}

	isParsingRules := true
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			isParsingRules = false
		}

		if isParsingRules {
			parseRule(rules, line)
		} else {
			res += checkLine(rules, line)
		}
	}

	fmt.Printf("Result: %d\n", res)
}

func main() {
	solvePuzzle01()
}
