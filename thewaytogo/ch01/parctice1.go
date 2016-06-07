package main

import (
	"fmt"
	"unicode/utf8"
)

type ByteSize float64

const (
	_ = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	ZB
)

func main() {
	var c1 complex64 = 5 + 10i
	fmt.Println(c1)

	var re float32 = 5.0
	var im float32 = 10.0

	c := complex(re, im)
	fmt.Println(c)

	fmt.Println(KB, MB, GB, TB, ZB)
}
