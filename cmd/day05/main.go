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

func sortLine(rules map[int][]int, line string) int {
	count := 0

	parsed := strings.Split(line, ",")
	data := make([]int, len(parsed))
	for idx := range parsed {
		data[idx], _ = strconv.Atoi(parsed[idx])
	}

	changed := true
mainloop:
	for changed {
		changed = false
		set := map[int]int{}

		for idx, val := range data {
			for _, val2 := range rules[val] {
				if idx2, ok := set[val2]; ok {
					val = val2
					tmp := data[idx]
					data[idx] = data[idx2]
					data[idx2] = tmp
					count++
					changed = true
					continue mainloop
				}
			}

			set[val] = idx
		}
	}

	if count == 0 {
		return 0
	}

	return data[len(data)/2]
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

func solvePuzzle02() {
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
			res += sortLine(rules, line)
		}
	}

	fmt.Printf("Result: %d\n", res)
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
