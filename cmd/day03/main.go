package main

import (
	"fmt"
	"strings"
)

func isChar(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || ch == '\''
}

func isNum(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func solvePuzzle01() {
	input := getInput()

	res := 0

	inst := ""
	paren := false
	comma := false
	a := 0
	b := 0

	for i := 0; i < len(input); i++ {
		reset := false

		if isChar(input[i]) {
			inst += string(input[i])
		} else if isNum(input[i]) {
			if inst == "mul" && paren {
				if !comma {
					a *= 10
					a += int(input[i] - '0')
				} else {
					b *= 10
					b += int(input[i] - '0')
				}
			} else {
				reset = true
			}
		} else {
			switch input[i] {
			case '(':
				if strings.HasSuffix(inst, "mul") {
					inst = "mul"
				}

				if inst == "mul" {
					paren = true
				} else {
					reset = true
				}
			case ')':
				reset = true
				if inst == "mul" && paren && comma {
					res += a * b
				}
			case ',':
				if inst == "mul" && paren {
					comma = true
				} else {
					reset = true
				}
			default:
				reset = true
			}
		}

		if reset {
			inst = ""
			paren = false
			comma = false
			a = 0
			b = 0
		}
	}

	fmt.Printf("Result: %d\n", res)
}

func solvePuzzle02() {
	input := getInput()

	res := 0

	inst := ""
	paren := false
	comma := false
	enabled := true
	a := 0
	b := 0

	for i := 0; i < len(input); i++ {
		reset := false

		if isChar(input[i]) {
			inst += string(input[i])
		} else if isNum(input[i]) {
			if enabled && inst == "mul" && paren {
				if !comma {
					a *= 10
					a += int(input[i] - '0')
				} else {
					b *= 10
					b += int(input[i] - '0')
				}
			} else {
				reset = true
			}
		} else {
			switch input[i] {
			case '(':
				if strings.HasSuffix(inst, "mul") {
					inst = "mul"
				}
				if strings.HasSuffix(inst, "do") {
					inst = "do"
				}
				if strings.HasSuffix(inst, "don't") {
					inst = "don't"
				}
				if inst == "mul" || inst == "do" || inst == "don't" {
					paren = true
				} else {
					reset = true
				}
			case ')':
				reset = true
				if enabled && inst == "mul" && paren && comma {
					res += a * b
				} else if inst == "do" && paren {
					enabled = true
				} else if inst == "don't" && paren {
					enabled = false
				}
			case ',':
				if inst == "mul" && paren {
					comma = true
				} else {
					reset = true
				}
			default:
				reset = true
			}
		}

		if reset {
			inst = ""
			paren = false
			comma = false
			a = 0
			b = 0
		}
	}

	fmt.Printf("Result: %d\n", res)
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
