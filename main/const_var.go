package main

import "fmt"

// 常量
const (
	A = iota
	B byte = 'b'
	C = 2
)

const (
	D = iota
	E = 3
	F = iota

	// 向下传导
	i = 1 << iota
	j
	k
)

// 变量
var (
	// 无符号整形
	ui8 uint = 8
	ui16 uint16 = 16
	ui32 uint32 = 32
	ui64 uint64 = 64

	// 有符号整形
	i8 int8 = 8 * -1
	i16 int16 = 16
	i32 int32 = 32 * -1
	i64 int64 = 64 * -1

	// 浮点
	f32 float32 = 3.5
	f64 float64 = 6.4

	// 布尔
	_bool_f bool = false
	_bool_t bool = true

	// 字节
	_byte byte = 0x10

	// 字符串
	s string = "y"
	z string = "z"
)

var v int

func main() {
	fmt.Println("const =", A, B, C, D, E, F, i, j, k, _bool_f, _bool_t)

	k := s + z
	fmt.Println(k)

	fmt.Println(s, z, s == z)

	g := s
	fmt.Println(g, s, g == s)

	var k1, k2, k3 int
	k1, k2, k3 = 1, 2, 3
	fmt.Print(k1, k2, k3)

	// 多变量赋值, '_'抛弃7
	v1, v2, v3, _ := 1, 3, 5, 7
	fmt.Println(v1, v2, v3)

	// 指针,指向变量地址
	var p *string
	p = &k
	fmt.Println(p)

	val := "val=原始值"
	fmt.Println(val)
	var q *string
	q = &val
	*q = "val=通过指针获取值之后做的修改"
	fmt.Println(val)

	var (
		aa int
		bb int
	)
	fmt.Println(aa, bb)
}