package main

import "unsafe"

const x, y int = 1, 2
const s = "Hello World"

const (
	a, b      = 10, 100
	c    bool = false
)

const (
	i = "abc"
	j
)

const (
	o = "abc"
	p = len(o)
	q = unsafe.Sizeof(a)
)

const (
	w byte = 100
	u int  = 1e20
)

func main() {
	const x = "xxx"

	println(j)

	println(o, p, q)

	println(w, u)
}
