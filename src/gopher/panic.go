package main

import (
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Runtime error caught: %v\n", r)
		}
	}()
	foo(0)
}

func foo(x int) {
	if x == 0 {
		panic("x doesn't be 0")
	}
}
