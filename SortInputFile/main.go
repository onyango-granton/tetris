package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Tetromino struct{
	shape [][]int
	name string
}

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

func main() {
	sampleFile, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// var lines []string
	var nums [][]int
	for i, ch := range strings.Split(string(sampleFile), "\n") {
		if ch == "" {
			continue
		}
		chArr, err := stringToIntSlice(ch)
		if err!= nil{
			fmt.Println(err.Error(),"at line",i+1)
			return
		} else {
			nums = append(nums, chArr)
		}
		// lines = append(lines, ch)
	}
	// fmt.Println(len(lines))

	// alphabetMap := make(map[string][][]int)

	// for i := 0; i < len(lines); {
	// 	for j := i; j < i+4; j++ {
	// 		fmt.Println(lines[j], i)
	// 	}
	// 	i += 4
	// }

	// for i:=0; i<len(nums);{

	// 	for j:=i;j<i+4;j++{
	// 		fmt.Println(lines[j],i)
	// 	}
	// 	i += 4
	// }
	
	startAscii := 'A'
	tetrominoes := make(map[rune][][]int)
	// for _,ch := range nums{
	// 	characterMino := [][]int{}
	// 	for range 4{
	// 		characterMino = append(characterMino, ch)
	// 	}
	// 	tetrominoes [startAscii] = characterMino
	// 	startAscii++
	// }


	for i:=0; i<len(nums);{
		characterMino := [][]int{}
		for j:=i; j<i+4;j++{
			characterMino = append(characterMino, nums[j])
		}
		tetrominoes[startAscii] = characterMino
		startAscii ++
		i += 4
	}

	for k,v := range tetrominoes{
		fmt.Println(string(k),v)
	}


	// for i := 0; i < len(nums);{
	// 	for j := i ; j < i+4;j++{
	// 		fmt.Print(nums[j])
	// 	}
	// 	fmt.Println(" ",i)
	// 	i += 4
	// }
}
