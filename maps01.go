package main

import (
	"fmt"
)

func main() {
	var isEnabled bool
	isEnabled = true
	m := map[string]string{}
	m["A"] = "a"
	if isEnabled {
		fmt.Println(m["A"])
	}
}
