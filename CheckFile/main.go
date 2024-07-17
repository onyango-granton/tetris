package main

import (
	"fmt"
)

func allOne(num1, num2 int) bool {
	if num1 == 1{
		return num1 == num2
	}
	return false
}

func isSurroundedByOnes(arr [][]int, row, col int) bool {
	// Check horizontally
	if col-1 >= 0 && allOne(arr[row][col-1],arr[row][col]) || col+1 < len(arr[row]) && allOne(arr[row][col+1],arr[row][col]) {
		return true
	}
	// Check vertically
	if row-1 >= 0 && allOne(arr[row][col],arr[row-1][col]) || row+1 < len(arr) && allOne(arr[row][col],arr[row+1][col]) {
		return true
	}
	return false
}

func isValidTetro(tetro [][]int) bool {
	var bordercount int
	var linecount int

	for row := 0; row < len(tetro); row++ {
		for col := 0; col < len(tetro[row]); col++ {
			if tetro[row][col] == 1{
				linecount++
			}
			if tetro[row][col] == 1 && isSurroundedByOnes(tetro, row, col) {
				bordercount++
				fmt.Printf("Element at (%d, %d) is surrounded by ones\n", row, col)
			}
		}
	}
	if bordercount == linecount{
		return true
	} else {
		return false
	}
}

func main() {
	arr := [][]int{
		{0, 1, 1, 1, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0},
	}

	var bordercount int
	var linecount int

	for row := 0; row < len(arr); row++ {
		for col := 0; col < len(arr[row]); col++ {
			if arr[row][col] == 1{
				linecount++
			}
			if arr[row][col] == 1 && isSurroundedByOnes(arr, row, col) {
				bordercount++
				fmt.Printf("Element at (%d, %d) is surrounded by ones\n", row, col)
			}
		}
	}

	if count > 4 || count2 > 4{
		fmt.Println("Invalid tetro")
	} else {
		fmt.Println("Valid tetro")
	}
}
