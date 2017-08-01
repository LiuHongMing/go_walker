/*
  接口：
      内嵌
      赋值
      任意类型
 */

package main

import (
	"fmt"
	"bufio"
)

// 类型
type Integer int

// 以值为接收者
//noinspection GoDuplicateFunctionOrMethod
func (a Integer) Less(b Integer) bool {
	return a < b
}

// 以指针为接收者
//noinspection GoDuplicateFunctionOrMethod
func (a *Integer) Add(b Integer) {
	*a += b
}

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

type Lesser interface {
	Less(b Integer) bool
}

// 内嵌
type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

func main() {
	var a Integer = 1

	// 以指针或值为接收者的区别在于：值方法可通过指针和值调用，而指针方法只能通过指针来调用。
	var b LessAdder = &a
	b.Add(11)
	println(b.Less(22))
	println(a)

	// 错误，不能实现以指针为接收者的Add方法
	// var b2 LessAdder = a

	var c Lesser = &a
	println(c.Less(22))

	var c2 Lesser = a
	println(c2.Less(22))

	// 任意类型
	var v1 interface{} = 1
	var v2 interface{} = "abc"
	var v3 interface{} = &v2
	var v4 interface{} = struct{ X int }{1}
	var v5 interface{} = &struct{ X int }{1}
	fmt.Println(v1, v2, v3, v4, v5)

}