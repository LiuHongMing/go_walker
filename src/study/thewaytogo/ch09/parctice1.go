package main

import (
	"container/list"
	"fmt"
	"unsafe"
)

func main() {

	l := list.New()
	l.PushBack(101)
	l.PushBack(102)
	l.PushBack(103)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	var b byte = 10
	var i int8 = 10
	fmt.Println(unsafe.Sizeof(b), unsafe.Sizeof(i))

}