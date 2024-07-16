package main

import (
	"fmt"
	"math"
)

// defining tetro structure
type Tetro struct {
	shape [][]int
	name  string
}

var (
	A = Tetro{[][]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {1, 1, 1, 1}}, "A"}
	B = Tetro{[][]int{{0, 0, 0, 0}, {0, 0, 0, 1}, {0, 1, 1, 1}, {0, 0, 0, 0}}, "B"}
	C = Tetro{[][]int{{1, 1, 0, 0}, {0, 1, 0, 0}, {0, 1, 0, 0}, {0, 0, 0, 0}}, "C"}
	D = Tetro{[][]int{{0, 0, 0, 0}, {0, 1, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}, "D"}
	E = Tetro{[][]int{{0, 0, 0, 0}, {0, 1, 1, 0}, {1, 1, 0, 0}, {0, 0, 0, 0}}, "E"}
	F = Tetro{[][]int{{0, 0, 0, 0}, {0, 0, 1, 0}, {0, 1, 1, 1}, {0, 0, 0, 0}}, "F"}
	G = Tetro{[][]int{{0, 0, 0, 0}, {1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}, "G"}
)

var (
	tetroGroup = []Tetro{A, B, C, D, E, F, G}
	gridSize   = math.Sqrt(float64(len(tetroGroup) * len(tetroGroup[0].shape)))
	grid       = make([][]string, int(math.Ceil(gridSize)))
)

// initializes a square grid 2D
func initGrid() {
	for i := 0; i < int(math.Ceil(gridSize)); i++ {
		grid[i] = make([]string, int(math.Ceil(gridSize)))
		for j := 0; j < int(math.Ceil(gridSize)); j++ {
			grid[i][j] = "*"
		}
	}
}

// checks whether we can place the term at the specific row and col of grid
func canPlace(term Tetro, grid [][]string, row, col int) bool {
	for r := range term.shape {
		for c := range term.shape[r] {
			if term.shape[r][c] == 1 {
				if row+r >= len(grid) || col+c >= len(grid[row]) || grid[row+r][col+c] != "*" {
					return false
				}
			}
		}
	}
	return true
}

// Places the term in row and col of the grid
func place(term Tetro, grid [][]string, row, col int) {
	for r := range term.shape {
		for c := range term.shape[r] {
			if term.shape[r][c] == 1 {
				grid[row+r][col+c] = term.name
			}
		}
	}
}

func remove(term Tetro, grid [][]string, row, col int) {
	for r := range term.shape {
		for c := range term.shape[r] {
			if term.shape[r][c] == 1 {
				grid[row+r][col+c] = "*"
			}
		}
	}
}
