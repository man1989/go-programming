package main

import "fmt"

func getCount() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	var getNextCount = getCount()
	fmt.Println(getNextCount())
	fmt.Println(getNextCount())
	fmt.Println(getNextCount())
}
