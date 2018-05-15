package main

import (
	"fmt"
	"reflect"
)

const some float64 = 1.2222

var (
	name   string
	course string
	module float64
)

func main() {
	fmt.Println("name type is ", reflect.TypeOf(name))
	fmt.Println("module type is ", reflect.TypeOf(module))
}
