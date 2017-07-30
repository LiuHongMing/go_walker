/*
  面向对象：
      类型
      值语义、引用语义
      结构体
 */

package main

import "fmt"

// 类型
type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

func (a *Integer) Add(b Integer) {
	*a += b
}

// 结构体
type Rect struct {
	x, y          float64
	width, height float64
}

func main() {
	var a Integer = 1
	if a.Less(2) {
		fmt.Println(a, "less 2")
	}
	a.Add(2)
	fmt.Println(a)

	// 值类型，值语义
	var x = [3]int{1, 2, 3}
	var y = x
	y[1]++
	fmt.Println(x, y)

	// 引用类型，引用语义
	var z = &x
	z[1]++
	fmt.Println(x, *z)

	rect1 := new(Rect)
	rect2 := &Rect{}
	rect3 := &Rect{0, 0, 100, 200}
	rect4 := &Rect{width: 100, height: 200}
	println(rect1, rect2, rect3, rect4)

	f := new(Foo)
	f.Age = 23
	f.Name = "Tom"
	f.Base.Name = "John"
	f.Foo()
	f.Bar()
	fmt.Println(f)
}

// 匿名组合
type Base struct {
	Name string
}

func (b *Base) Foo() {
	fmt.Println("Base.Foo()")
}

func (b *Base) Bar() {
	fmt.Println("Base.Bar()")
}

type Foo struct {
	Name string
	Age int
	Base
}

func (f *Foo) Bar() {
	f.Base.Bar()
}
