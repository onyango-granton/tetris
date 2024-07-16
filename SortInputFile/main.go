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

func stringToIntSlice(s string) ([]int,error){
	res := []int{}
	if len(s) != 4{
		return nil, errors.New("invalid length entry in file")
	}
	for _,b := range s{
		num, err := byteToInt(byte(b))
		if err != nil{
			return nil, err
		}
		res = append(res, num)
	}
	return res, nil
}
