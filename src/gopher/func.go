/*
  函数
 */

package main

import (
	"gopher/mymath"
	"fmt"
)

func main() {
	v, err := mymath.Add(-1, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}

	// 匿名函数
	f := func(x, y int) int {
		return x + y
	}
	f(5, 8)

	// 匿名函数调用
	//func(ch chan int) {
	//	ch <- 1
	//}(make(chan int))
}
