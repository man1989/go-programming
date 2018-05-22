package main

import "fmt"

func main() {
	var str = make([]string, 5)
	c := 'a'
	l := 0
	for i := 0; i < 10; i++ {
		l = int(c) + i
		str[i] = string(l)
	}
	fmt.Println(str, len(str), cap(str))
}
