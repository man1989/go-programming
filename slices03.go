package main

import (
	"fmt"
)

func main() {
	slices := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	a := slices[2:5]
	printSlices(a)
	a[1] = 10
	fmt.Println("original slice", slices)
	printSlices(a)
	r := a[:2]
	r[0] = 25
	fmt.Printf("original %d, slice %d, re-slice %d\n", slices, a, r)
}

func printSlices(s []int) {
	fmt.Printf("len %d, cap %d, %d\n", len(s), cap(s), s)
}
