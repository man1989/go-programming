package main

import "fmt"

func main() {
	max := getMax(1, 2, 3, 41, 5, 6)
	fmt.Println("max is: ", max)
}

func getMax(list ...int) int {
	var max = list[0]
	for _, next := range list[1:] {
		if max < next {
			max = next
		}
	}
	return max
}
