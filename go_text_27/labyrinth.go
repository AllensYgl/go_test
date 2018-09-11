package main

import (
	"fmt"
	"os"
)

//read the labyrinthMap
func readLabyrinth(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	labyrinth := make([][]int, row)
	for i := range labyrinth {
		labyrinth[i] = make([]int, col)
		for j := range labyrinth[i] {
			fmt.Fscanf(file, "%d", &labyrinth[i][j])
		}
	}
	return labyrinth
}

type point struct {
	i, j int
}

var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) || p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true

}

func walk(labyrinth [][]int, start, end point) [][]int {
	steps := make([][]int, len(labyrinth))
	for i := range steps {
		steps[i] = make([]int, len(labyrinth[i]))
	}

	quenue := []point{start}

	for len(quenue) > 0 {
		cur := quenue[0]
		quenue := quenue[1:]

		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)

			//next is 0
			//steps is 0
			//is !=start
			if value, ok := next.at(labyrinth); !ok || value == 1 {
				continue
			}
			if value, ok := next.at(steps); value != 0 || !ok {
				continue
			}

			if next == start {
				continue
			}
			cursteps, _ := cur.at(steps)
			steps[next.i][next.j] = cursteps + 1

			quenue = append(quenue, next)
		}
	}
	return steps
}

func main() {
	labyrinth := readLabyrinth("labyrinth.txt")

	step := walk(labyrinth, point{0, 0}, point{len(labyrinth) - 1, len(labyrinth[0]) - 1})

	for _, row := range step {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}
