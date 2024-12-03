package main

import "fmt"

func isChar(ch byte) bool {
	return ch >= 'a' && ch <= 'z'
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

func main() {
	solvePuzzle01()
}
