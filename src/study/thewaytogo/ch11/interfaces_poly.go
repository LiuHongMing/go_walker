package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Squre struct {
	side float32
}

type Rectangle struct {
	length, width float32
}

func (sq *Squre) Area() float32 {
	return sq.side * sq.side
}

func (re Rectangle) Area() float32 {
	return re.length * re.width
}

func main() {
	re := Rectangle{5, 3}
	sq := &Squre{5}

	shapers := []Shaper{re, sq}
	for n, _ := range shapers {
		fmt.Println("Shape details:", shapers[n])
		fmt.Println("Area of this shape is", shapers[n].Area())
	}
}
