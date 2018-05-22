package main

import "fmt"

func main() {
	var s1 []int
	printSlices(s1)
	var s2 = make([]int, 5)
	printSlices(s2)
	s3 := s2[2:4] //slice
	printSlices(s3)
	s4 := s3[:3]
	printSlices(s4)
}

func printSlices(s []int) {
	fmt.Printf("length = %d, capacity = %d, %d\n", len(s), cap(s), s)
}
