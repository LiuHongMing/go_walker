/*
  指针
 */

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	type data struct {
		a int
	}

	var d = data{1234}
	var p *data

	p = &d

	fmt.Printf("%p, %v\n", p, p.a)

	x := 0x12345679
	q := unsafe.Pointer(&x)
	n := (*[4]byte)(q)

	for i := 0; i < len(n); i++ {
		fmt.Printf("%X ", n[i])
	}

	y := struct {
		s string
		x int
	}{"abc", 100}

	w := uintptr(unsafe.Pointer(&y))
	w += unsafe.Offsetof(y.x)

	w2 := unsafe.Pointer(w)
	px := (*int) (w2)
	*px = 200

	fmt.Printf("%#v\n", y)
}


func test3() *int {
	x := 100
	return &x
}
