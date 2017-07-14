/**
学习变量的类型、声明
 */

package main

var x, y, z int
var s, n = "abc", 123

var (
	a int
	b float32
)

func main() {
	x := 3 // 局部变量声明、赋值
	println(x)

	//s, n = 0x1234, "Hello World!" // 已声明类型无法改变
	//println(x, s, n)

	data, i := [3]int{1, 2, 3}, 0
	i, data[i] = 2, 100

	println(data[i], data[0])
}
