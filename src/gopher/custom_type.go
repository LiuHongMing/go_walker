/*
  自定义类型
 */

package main

var a struct{ x int `a` }
var b struct{ x int `ab` }

func main() {
	// println(b == a) // invalid operation: b == a (mismatched types struct { x int "ab" } and struct { x int "a" })

	type bigint int64
	var x bigint = 100
	println(x)

	x = 1234
	var b = bigint(x)
	var b2 = int64(b)
	println(b2)
}
