package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseInput(input string) []int {
	var res []int

	numbers := strings.Split(input, " ")
	for _, number := range numbers {
		num, _ := strconv.Atoi(number)
		res = append(res, num)
	}

	return res
}

func countDigits(num int) int {
	res := 0

	for num > 0 {
		res++
		num /= 10
	}

	return res
}

func splitNumber(num int, middle int) (int, int) {
	digits := make([]int, middle)
	idx := 0

	for middle > 0 {
		middle--
		digits[idx] = num % 10
		idx++
		num /= 10
	}

	num2 := 0
	idx--
	for idx >= 0 {
		num2 *= 10
		num2 += digits[idx]
		idx--
	}

	return num, num2
}

func numOfStones(num, n int, mem map[string]int) int {
	if n == 0 {
		return 1
	}

	key := fmt.Sprintf("%d-%d", num, n)
	if val, ok := mem[key]; ok {
		return val
	}

	res := 0

	if num == 0 {
		res = numOfStones(1, n-1, mem)
	} else {
		count := countDigits(num)
		if count%2 == 0 {
			n1, n2 := splitNumber(num, count/2)
			res += numOfStones(n1, n-1, mem)
			res += numOfStones(n2, n-1, mem)
		} else {
			res = numOfStones(num*2024, n-1, mem)
		}
	}

	mem[key] = res

	return res
}

func solvePuzzle01() {
	input := getInput()
	stones := parseInput(input)

	res := 0

	mem := map[string]int{}
	for _, stone := range stones {
		res += numOfStones(stone, 25, mem)
	}

	fmt.Printf("Number of stones: %d\n", res)
}

func solvePuzzle02() {
	input := getInput()
	stones := parseInput(input)

	res := 0

	mem := map[string]int{}
	for _, stone := range stones {
		res += numOfStones(stone, 75, mem)
	}

	fmt.Printf("Number of stones: %d\n", res)
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
