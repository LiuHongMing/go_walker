package main

import "fmt"

type Any interface{}

type Empty interface{}

type EmptyImpl struct {
	name string
}

func testEmpty(any Any) {
	if p, ok := any.(EmptyImpl); ok {
		fmt.Println("接口:", p.name)
	}
	fmt.Println("空接口是顶层接口")
}


func main() {
	var e Empty = EmptyImpl{"空接口"}
	testEmpty(e)
}
