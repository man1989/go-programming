package main

import (
	"fmt"
	"os"
)

func main() {
	var env = os.Environ()
	for _, env := range env {
		fmt.Println(env)
	}
	fmt.Println(os.Getenv("PATH"))
}
