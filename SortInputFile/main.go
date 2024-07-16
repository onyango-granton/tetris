package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func byteToInt(b byte) (int, error) {
	if b == '.' {
		return 0, nil
	} else if b == '#' {
		return 1, nil
	} else {
		return 0, errors.New("error in sample file")
	}
}

func isValidTetro(tetro [][]int) bool {
	for row := range tetro{
		for col := range tetro[row]{
			if tetro[row][col] == 1 && tetro
		}
	}
}

