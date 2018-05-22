package main

import "fmt"

func main() {
	name := "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	// name[0] = "A"
	// change(&name)
	fmt.Println(name)
}

func change(name *string) {
	*name = "x"
	fmt.Println(*name)
}
