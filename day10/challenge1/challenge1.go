package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	row int
	col int
}

type Board struct {
	trailheads []Coord
	board      [][]int
	visited    [][]bool
}

func NewBoard(trailheads []Coord, board [][]int) *Board {
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[i]))
	}
	return &Board{trailheads: trailheads, board: board, visited: visited}
}

func (b *Board) coordIsInBoard(c Coord) bool {
	if c.row < 0 || c.col < 0 {
		return false
	}
	if c.row >= len(b.board) || c.col >= len(b.board[c.row]) {
		return false
	}
	return true
}
func (b *Board) getNeighbours(c Coord) []Coord {
	var neighbours []Coord
	for _, d := range [][2]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}} {
		if b.coordIsInBoard(Coord{c.row + d[0], c.col + d[1]}) {
			neighbours = append(neighbours, Coord{c.row + d[0], c.col + d[1]})
		}
	}
	return neighbours
}

func (b *Board) traverse(c Coord, val int) {
	if val == 9 {
		b.visited[c.row][c.col] = true
	}
	neighbours := b.getNeighbours(c)
	for _, n := range neighbours {
		if b.board[n.row][n.col] == val+1 {
			b.traverse(n, val+1)
		}
	}
}

type Path struct {
	coord    Coord
	children []Path
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input [][]int
	var trailheads []Coord
	var trailends []Coord
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		linea := strings.Split(line, "")
		var vals []int

		for col, v := range linea {
			if v == "." {
				vals = append(vals, -1)
			} else {
				i, err := strconv.Atoi(v)
				if err != nil {
					panic("Error parsing input")
				}
				vals = append(vals, i)
				if i == 0 {
					trailheads = append(trailheads, Coord{row: row, col: col})
				}
				if i == 9 {
					trailends = append(trailends, Coord{row: row, col: col})
				}
			}
		}
		input = append(input, vals)
		row += 1
	}
	fmt.Println(input)
	fmt.Println(trailheads)
	//board := NewBoard(trailheads, input)
	fmt.Println("Part 1")
	count := 0
	for _, t := range trailheads {
		fmt.Println("traversing from", t)
		board := NewBoard(trailheads, input)

		board.traverse(t, 0)
		for _, t := range trailends {
			if board.visited[t.row][t.col] {
				count += 1
			}
		}
	}

	fmt.Println("done")
	fmt.Println(count)

}
