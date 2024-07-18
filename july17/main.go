package main

import (
    "fmt"
)

type Tetromino struct {
    shape [][]int
    name  string
}

var A = Tetromino{[][]int{{0, 0, 0, 1}, {0, 0, 0, 1}, {0, 0, 0, 1}, {0, 0, 0, 1}}, "A"}
var B = Tetromino{[][]int{ {1, 1, 1, 1}}, "B"}
var C = Tetromino{[][]int{{0, 1, 1, 1}, {0, 0, 0, 1},}, "C"}
var D = Tetromino{[][]int{{0, 0, 1, 1}, {0, 1, 1, 0}}, "D"}
var E = Tetromino{[][]int{{0, 1, 1, 0}, {0, 1, 1, 0}}, "E"}
var F = Tetromino{[][]int{{1, 1, 0, 0}, {0, 1, 1, 0}}, "F"}
var G = Tetromino{[][]int{{1, 1, 0, 0}, {0, 1, 0, 0}, {0, 1, 0, 0}}, "G"}
var H = Tetromino{[][]int{{1, 1, 1, 0}, {0, 1, 0, 0}}, "H"}

var tetrominoes = []Tetromino{A, B, C, D, E, F, G, H}

var gridSize = 6
var grid = make([][]string, gridSize)

func initGrid() {
    for i := range grid {
        grid[i] = make([]string, gridSize)
        for j := range grid[i] {
            grid[i][j] = "*"
        }
    }
}

func canPlace(tetromino Tetromino, grid [][]string, row, col int) bool {
    shape := tetromino.shape
    for r := range shape {
        for c := range shape[0] {
            if shape[r][c] == 1 {
                if row+r >= len(grid) || col+c >= len(grid[0]) || grid[row+r][col+c] != "*" {
                    return false
                }
            }
        }
    }
    return true
}

func place(tetromino Tetromino, grid [][]string, row, col int) {
    shape := tetromino.shape
    for r := range shape {
        for c := range shape[0] {
            if shape[r][c] == 1 {
                grid[row+r][col+c] = tetromino.name
            }
        }
    }
}

func remove(tetromino Tetromino, grid [][]string, row, col int) {
    shape := tetromino.shape
    for r := range shape {
        for c := range shape[0] {
            if shape[r][c] == 1 {
                grid[row+r][col+c] = "*"
            }
        }
    }
}

func solve(tetrominoes []Tetromino, grid [][]string, index int) bool {
    if index == len(tetrominoes) {
        return true
    }
    for row := 0; row < len(grid); row++ {
        for col := 0; col < len(grid[0]); col++ {
            if canPlace(tetrominoes[index], grid, row, col) {
                place(tetrominoes[index], grid, row, col)
                if solve(tetrominoes, grid, index+1) {
                    return true
                }
                remove(tetrominoes[index], grid, row, col)
            }
        }
    }
    return false
}

func printGrid(grid [][]string) {
    for _, row := range grid {
        for _, cell := range row {
            fmt.Print(cell, " ")
        }
        fmt.Println()
    }
}

func main() {
    initGrid()
    if solve(tetrominoes, grid, 0) {
        printGrid(grid)
    } else {
        fmt.Println("No solution found")
    }
}
