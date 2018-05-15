package main

import "fmt"
import "os"

func sayHello() {
	fmt.Println("Hello, ", os.Getenv("USER"))
}

func getSum(x, y int) int {
	return x + y
}

func oddEven(n []int) ([]int, []int) {
	var x []int
	var y []int
	for i := 0; i < len(n); i++ {
		if n[i]%2 == 0 {
			x = append(x, n[i])
		} else {
			y = append(y, n[i])
		}
	}
	return x, y
}

//named return
func splitInt(n []int) (x, y int) {
	x = n[:1][0]
	y = n[2:][1]
	return
}

func main() {
	sayHello()
	fmt.Println(getSum(5, 5))
	fmt.Println(splitInt([]int{1, 2, 3, 4, 5}))
	fmt.Println(oddEven([]int{1, 2, 3, 4, 5, 6}))
}
