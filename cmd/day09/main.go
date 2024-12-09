package main

import "fmt"

func parseInput(input string) []int {
	var res []int

	id := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < int(input[i]-'0'); j++ {
			if i%2 == 0 {
				res = append(res, id)
			} else {
				res = append(res, -1)
			}
		}

		if i%2 == 0 {
			id++
		}
	}

	return res
}

func rearrange(data []int) {
	l := 0
	r := len(data) - 1

	for {
		for data[l] >= 0 {
			l++
		}
		for data[r] < 0 {
			r--
		}

		if l >= r {
			break
		}

		data[l] = data[r]
		data[r] = -1
	}
}

func checksum(data []int) int {
	res := 0

	for i := range data {
		if data[i] < 0 {
			break
		}

		res += data[i] * i
	}

	return res
}

func solvePuzzle01() {
	input := getInput()
	data := parseInput(input)

	rearrange(data)
	res := checksum(data)

	fmt.Printf("Checksum: %d\n", res)
}

func main() {
	solvePuzzle01()
}
