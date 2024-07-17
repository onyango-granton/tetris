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

// func isValidTetro(tetro [][]int) bool {
// 	for row := range tetro{
// 		for col := range tetro[row]{
// 			if tetro[row][col] == 1 && tetro
// 		}
// 	}
// }

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

func isValidTetro(tetro [][]int) (bool,error) {
	var bordercount int
	var linecount int

	for row := 0; row < len(tetro); row++ {
		for col := 0; col < len(tetro[row]); col++ {
			if tetro[row][col] == 1{
				linecount++
			}
			if tetro[row][col] == 1 && isSurroundedByOnes(tetro, row, col) {
				bordercount++
				// fmt.Printf("Element at (%d, %d) is surrounded by ones\n", row, col)
			}
		}
	}
	if bordercount > 4 || linecount > 4{
		return false, errors.New("Invalid Tetromino")
	} else {
		return true, nil
	}
}


func tetroGroup(textFile string) []Tetromino {
	tetrominoesGroup := []Tetromino{}
	// opens text file
	sampleFile, err := os.ReadFile(textFile)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	var nums [][]int
	for i, ch := range strings.Split(string(sampleFile), "\n") {
		if ch == "" {
			continue
		}
		chArr, err := stringToIntSlice(ch)
		if err!= nil{
			fmt.Println(err.Error(),"at line",i+1)
			return nil
		} else {
			nums = append(nums, chArr)
		}
	}

	startAscii := 'A'
	tetrominoes := make(map[rune][][]int)

	for i:=0; i<len(nums);{
		characterMino := [][]int{}
		for j:=i; j<i+4;j++{
			characterMino = append(characterMino, nums[j])
		}
		tetrominoes[startAscii] = characterMino
		startAscii ++
		i += 4
	}

	for k,_ := range tetrominoes{
		res, err := isValidTetro(tetrominoes[k])
		if err != nil{
			fmt.Println(err.Error())
			return nil
		} else if res {
			newTetro := Tetromino{shape: tetrominoes[k], name: string(k)}
			tetrominoesGroup = append(tetrominoesGroup, newTetro)
		}
	}

	return tetrominoesGroup

	// for k,_ := range tetrominoes{
	// 	newTetro := Tetromino{shape: tetrominoes[k], name: string(k)}
	// 	tetrominoesGroup = append(tetrominoesGroup, newTetro)
	// }
}





func main() {
	sampleFile, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// var lines []string
	tetrominoesGroup := []Tetromino{}
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

	// for k,_ := range tetrominoes{
	// 	if isValidTetro(tetrominoes[k]){
	// 		fmt.Println(true)
	// 	} else {
	// 		fmt.Println(false)
	// 	}
	// }

	for k,_ := range tetrominoes{
		newTetro := Tetromino{shape: tetrominoes[k], name: string(k)}
		tetrominoesGroup = append(tetrominoesGroup, newTetro)
	}


	// for i := 0; i < len(nums);{
	// 	for j := i ; j < i+4;j++{
	// 		fmt.Print(nums[j])
	// 	}
	// 	fmt.Println(" ",i)
	// 	i += 4
	// }
}
