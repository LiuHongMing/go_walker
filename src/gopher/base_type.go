/*
  基本类型
  类型转换
 */

package main

import "math"

func main() {
	a, b, c, d := 071, 0x1F, 1e9, math.MinInt16
	println(a, b, c, d)

	var m byte = 100
	var n int = int(m)
	println(n)

	if n { // non-bool n (type int) used as if condition
		println("true")
	}
}
