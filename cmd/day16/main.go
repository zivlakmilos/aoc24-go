package main

import (
	"container/heap"
	"fmt"
	"strings"
)

const Inf = ^uint(0x00)

type Cell struct {
	score   uint
	value   byte
	visited bool
}

type Point struct {
	x int
	y int
}

type Node struct {
	pos Point
	o   Point
}

var neighbours = []Point{
	{
		x: 1,
		y: 0,
	},
	{
		x: -1,
		y: 0,
	},
	{
		x: 0,
		y: 1,
	},
	{
		x: 0,
		y: -1,
	},
}

func parseInput(input string) [][]Cell {
	var res [][]Cell

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		var row []Cell
		for i := range line {
			row = append(row, Cell{
				value: line[i],
				score: Inf,
			})
		}
		res = append(res, row)
	}

	return res
}

func foundCell(maze [][]Cell, ch byte) (int, int) {
	for y := range maze {
		for x := range maze[y] {
			if maze[y][x].value == ch {
				return x, y
			}
		}
	}

	return 0, 0
}

func solveMaze(start Point, end Point, maze [][]Cell) (uint, []Point) {
	dist := make([][]uint, len(maze))
	prev := make([][]Node, len(maze))

	for i := range maze {
		dist[i] = make([]uint, len(maze[i]))
		prev[i] = make([]Node, len(maze[i]))

		for j := range maze[i] {
			dist[i][j] = Inf
			prev[i][j] = Node{
				pos: Point{
					x: -1,
					y: -1,
				},
				o: Point{
					x: -1,
					y: -1,
				},
			}
		}
	}

	pq := &PriorityQueue[Node]{}
	heap.Init(pq)
	heap.Push(pq, &Item[Node]{
		value: Node{
			pos: start,
			o: Point{
				x: 1,
				y: 0,
			},
		},
		priority: 0,
	})
	dist[start.y][start.x] = 0

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item[Node])
		curPoint := current.value.pos
		curO := current.value.o

		for _, n := range neighbours {
			nx := curPoint.x + n.x
			ny := curPoint.y + n.y

			if maze[ny][nx].value == '#' {
				continue
			}

			nDist := dist[curPoint.y][curPoint.x] + 1
			if curO != n {
				nDist += 1000
			}

			if nDist < dist[ny][nx] {
				dist[ny][nx] = nDist
				prev[ny][nx] = current.value
				heap.Push(pq, &Item[Node]{
					value: Node{
						pos: Point{
							x: nx,
							y: ny,
						},
						o: n,
					},
					priority: int(nDist),
				})
			}
		}
	}

	var path []Point
	cur := end
	for cur != start {
		path = append([]Point{cur}, path...)
		cur = prev[cur.y][cur.x].pos
	}
	path = append([]Point{start}, path...)

	return dist[end.y][end.x], path
}

func solvePuzzle01() {
	input := getInput()
	maze := parseInput(input)
	sx, sy := foundCell(maze, 'S')
	ex, ey := foundCell(maze, 'E')

	start := Point{
		x: sx,
		y: sy,
	}
	end := Point{
		x: ex,
		y: ey,
	}

	res, _ := solveMaze(start, end, maze)

	fmt.Printf("Result: %d\n", res)
}

func main() {
	solvePuzzle01()
}
