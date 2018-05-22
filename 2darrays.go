package main

import (
	"fmt"
)

func main() {
	var arr [10]int
	var _2dArr [5][5]int
	// arr = append(arr, 1, 2, 3)
	fmt.Println(arr)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			_2dArr[i][j] = i + j
		}
	}
	fmt.Println("2d array: ", _2dArr)
}
