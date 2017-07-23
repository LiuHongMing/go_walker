/*
  基本类型
  类型转换
 */

package main

func main() {

	var v1 int32 = 10
	v2 := int32(v1) // 默认int
	println(v1, v2)

	var fv1 float32
	fv1 = 12
	fv2 := float32(12.0) // 默认float64
	println(fv1 == fv2)

	var val1 complex64
	val1 = 3.2 + 12i
	val2 := 3.2 + 12i // 默认complex128
	val3 := complex(3.2, 12)
	println(val1, val2, val3, real(val3), imag(val3))

}
