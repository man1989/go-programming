package main

import (
	"fmt"
)

var (
	name   = "Manish"
	course string
	module float64
)

func main() {
	var sum func(float64, float64) float64
	sum = func(a float64, b float64) float64 {
		return a + b
	}
	fmt.Println("Name is : ", name)
	fmt.Println("Course is : ", course)
	fmt.Println("moudle currently taking", module)
	fmt.Println(sum(2, 5));
}
