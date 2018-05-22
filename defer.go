package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("starting the countdown: ")
	for i := 10; i >= 0; i-- {
		time.Sleep(time.Second * 1)
		fmt.Println(i)
	}
	fmt.Println("Boom!!!")

}
