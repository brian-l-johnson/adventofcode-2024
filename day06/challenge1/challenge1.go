package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Coord struct {
	row int
	col int
}

type FloorMap struct {
	floor     [][]string
	visited   [][]bool
	guard_loc Coord
	guard_dir int
	direction [4][2]int
}

func (m *FloorMap) getNextGuardCoord() Coord {
	return Coord{row: m.guard_loc.row + m.direction[m.guard_dir][0], col: m.guard_loc.col + m.direction[m.guard_dir][1]}
}

func (m *FloorMap) guardIsObstructed() bool {
	nextCoord := m.getNextGuardCoord()
	if nextCoord.row < 0 || nextCoord.col < 0 || nextCoord.row >= len(m.floor) || nextCoord.col >= len(m.floor[nextCoord.row]) {
		return false
	}
	if m.floor[m.guard_loc.row+m.direction[m.guard_dir][0]][m.guard_loc.col+m.direction[m.guard_dir][1]] == "#" {
		return true
	}
	return false
}

func (m *FloorMap) print() {
	for r, _ := range m.floor {
		for c, _ := range m.floor[r] {
			if m.visited[r][c] {
				fmt.Print("X")
			} else {
				fmt.Print(m.floor[r][c])
			}
		}
		fmt.Println()
	}
}

func (m *FloorMap) guardWalk() {
	//walk the floor until the guard goes off the end of the map
	for m.guard_loc.row >= 0 && m.guard_loc.col >= 0 && m.guard_loc.row < len(m.floor) && m.guard_loc.col < len(m.floor[m.guard_loc.row]) {
		if m.guardIsObstructed() {
			m.guard_dir = (m.guard_dir + 1) % 4
		} else {
			m.visited[m.guard_loc.row][m.guard_loc.col] = true
			m.guard_loc.row += m.direction[m.guard_dir][0]
			m.guard_loc.col += m.direction[m.guard_dir][1]
		}
	}
	fmt.Println("done walking")
	sum := 0
	for r := 0; r < len(m.visited); r++ {
		for c := 0; c < len(m.visited[r]); c++ {
			if m.visited[r][c] {
				sum++
			}
		}
	}
	fmt.Println("visited squares: ", sum)

}

// instantiate a new FloorMap from a two dimensional array of strings
func NewFloorMap(floor [][]string) *FloorMap {
	//create a two dimenstional array of bools with the same dimensions as the floor map
	var visited [][]bool
	for i := 0; i < len(floor); i++ {
		row := make([]bool, len(floor[i]))
		for i := 0; i < len(floor[0]); i++ {
			row[i] = false
		}
		visited = append(visited, make([]bool, len(floor[i])))
	}

	guard_loc := Coord{-1, -1}
	for i := 0; i < len(floor); i++ {
		for j := 0; j < len(floor[i]); j++ {
			if floor[i][j] == "^" {
				guard_loc = Coord{row: i, col: j}
				fmt.Println("Guard location: ", guard_loc)
				break
			}
		}
	}
	direction := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	return &FloorMap{
		floor:     floor,
		visited:   visited,
		guard_loc: guard_loc,
		direction: direction,
		guard_dir: 0,
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var floormap [][]string

	for scanner.Scan() {
		line := scanner.Text()
		floormap = append(floormap, strings.Split(line, ""))
	}
	//fmt.Println(floormap)
	m := NewFloorMap(floormap)
	//m.print()
	//fmt.Println(m.visited)
	m.guardWalk()
	m.print()
	sum := 0
	for r := range m.floor {
		for c, _ := range m.floor[r] {
			if m.visited[r][c] {
				sum++
			}
		}
	}
	fmt.Println("Visited squares: ", sum)

}
