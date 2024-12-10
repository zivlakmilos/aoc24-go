package main

import (
	"fmt"
	"slices"
)

type Block struct {
	pos int
	len int
	id  int
}

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
			continue
		}

		res += data[i] * i
	}

	return res
}

func extractBlocks(data []int) ([]Block, []Block) {
	var memBlocks []Block
	var freeBlocks []Block
	pos := 0

	for pos < len(data) {
		size := 0
		for r := pos; r < len(data) && data[pos] == data[r]; r++ {
			size++
		}

		if data[pos] < 0 {
			freeBlocks = append(freeBlocks, Block{
				pos: pos,
				len: size,
				id:  data[pos],
			})
		} else {
			memBlocks = append(memBlocks, Block{
				pos: pos,
				len: size,
				id:  data[pos],
			})
		}

		pos += size
	}

	return memBlocks, freeBlocks
}

func moveBlock(data []int, t, f, l int) {
	for i := 0; i < l; i++ {
		data[t+i] = data[f+i]
		data[f+i] = -1
	}
}

func defragment(data []int) {
	memBlocks, freeBlocks := extractBlocks(data)

	l := 0
	for l < len(freeBlocks) && len(memBlocks) > 0 {
		r := len(memBlocks) - 1
		for r >= 0 && (memBlocks[r].len > freeBlocks[l].len) {
			r--
		}
		if r < 0 || memBlocks[r].pos < freeBlocks[l].pos {
			l++
			continue
		}

		moveBlock(data, freeBlocks[l].pos, memBlocks[r].pos, memBlocks[r].len)

		diff := freeBlocks[l].len - memBlocks[r].len
		freeBlocks[l].pos += freeBlocks[l].len - diff
		freeBlocks[l].len = diff
		if freeBlocks[l].len <= 0 {
			l++
		}

		memBlocks = slices.Delete(memBlocks, r, r+1)
	}
}

func solvePuzzle01() {
	input := getInput()
	data := parseInput(input)

	rearrange(data)
	res := checksum(data)

	fmt.Printf("Checksum: %d\n", res)
}

func solvePuzzle02() {
	input := getInput()
	data := parseInput(input)

	defragment(data)
	res := checksum(data)

	fmt.Printf("Checksum: %d\n", res)
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
