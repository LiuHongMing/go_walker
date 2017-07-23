/**
  枚举
 */

package main

const (
	Sunday    = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	numberOfDays // 这个常量没有导出
)

const (
	_        = 0
	KB int64 = 1 << (10 * iota)
	MB
	GB
	TB
)

const (
	A, B = iota, iota << 10
	C, D
)

const (
	X = iota
	Y
	Z = "Z"
	U
	V = iota // 显示恢复, 包含上面两行
	W
)

type Color int

const (
	Black Color = iota
	Red
	Blue
)

func test2(c Color) {}

func main() {
	println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, numberOfDays)
	println(KB, MB, GB, TB)
	println(A, B, C, D)
	println(X, Y, Z, U, V, W)

	c := Black
	test2(c)

	//x := 1
	//test(x) // 类型匹配错误

	test2(1) // 常量被编译器自动转换
}
