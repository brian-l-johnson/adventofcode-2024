package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Coord struct {
	row int
	col int
}

type FloorMap struct {
	floor     [][]string
	visited   [][][]int
	guard_loc Coord
	guard_dir int
	direction [4][2]int
}

func (m *FloorMap) getVisitedCoords() []Coord {
	var coords []Coord
	for r, _ := range m.floor {
		for c, _ := range m.floor[r] {
			if len(m.visited[r][c]) > 0 {
				coords = append(coords, Coord{row: r, col: c})
			}
		}
	}
	return coords
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
			if len(m.visited[r][c]) > 0 {
				fmt.Print("X")
			} else {
				fmt.Print(m.floor[r][c])
			}
		}
		fmt.Println()
	}
}

func (m *FloorMap) guardWalk() int {
	//walk the floor until the guard goes off the end of the map
	for m.guard_loc.row >= 0 && m.guard_loc.col >= 0 && m.guard_loc.row < len(m.floor) && m.guard_loc.col < len(m.floor[m.guard_loc.row]) {
		if m.guardIsObstructed() {
			m.guard_dir = (m.guard_dir + 1) % 4
		} else {
			if slices.Contains(m.visited[m.guard_loc.row][m.guard_loc.col], m.guard_dir) {
				return 1
			}
			m.visited[m.guard_loc.row][m.guard_loc.col] = append(m.visited[m.guard_loc.row][m.guard_loc.col], m.guard_dir)
			m.guard_loc.row += m.direction[m.guard_dir][0]
			m.guard_loc.col += m.direction[m.guard_dir][1]
		}
	}
	//fmt.Println("done walking")
	sum := 0
	for r := 0; r < len(m.visited); r++ {
		for c := 0; c < len(m.visited[r]); c++ {
			if len(m.visited[r][c]) > 0 {
				sum++
			}
		}
	}
	//fmt.Println("visited squares: ", sum)
	return 0

}

// instantiate a new FloorMap from a two dimensional array of strings
func NewFloorMap(floor [][]string) *FloorMap {
	//create a two dimenstional array of slices to keep track of visited coodinates and directions

	var visited [][][]int
	for i := 0; i < len(floor); i++ {
		visited = append(visited, make([][]int, len(floor[i])))
	}

	guard_loc := Coord{-1, -1}
	for i := 0; i < len(floor); i++ {
		for j := 0; j < len(floor[i]); j++ {
			if floor[i][j] == "^" {
				guard_loc = Coord{row: i, col: j}
				//fmt.Println("Guard location: ", guard_loc)
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

	visited := m.getVisitedCoords()
	//fmt.Printf("visited %d squares:  %v\n", len(visited), visited)

	loops := 0
	/*
		var visited []Coord
		for r, _ := range floormap {
			for c, _ := range floormap[r] {
				visited = append(visited, Coord{row: r, col: c})
			}
		}
	*/

	for _, v := range visited {
		//fmt.Println(v)
		var modifiedFloorMap [][]string
		for _, row := range floormap {
			modifiedFloorMap = append(modifiedFloorMap, slices.Clone(row))
		}
		modifiedFloorMap[v.row][v.col] = "#"
		mm := NewFloorMap(modifiedFloorMap)
		if mm.guardWalk() == 1 {
			//fmt.Println("loop found for: ", v)
			loops++
		}
	}
	fmt.Println("loops found:  ", loops)

}
