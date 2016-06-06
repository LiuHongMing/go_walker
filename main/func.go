package main

import "fmt"

func main() {
	// 函数调用
	a := 1
	v1 := A(a)
	fmt.Println(v1)

	// 多值返回
	b, c, d := 4, 5, 6
	v1, v2 := B(b, c, d)
	fmt.Println(v1, v2)

	// 函数作为参数
	C(A, Ok)

	// 匿名函数
	nano := func() {
		fmt.Println("nano")
	}
	nano()

	// 闭包
	f := closure(10)
	fmt.Println(f(1))
	fmt.Println(f(5))

	// defer, 按照LIFO的顺序执行
	defer fmt.Println("defer_1")
	defer fmt.Println("defer_2")
	defer fmt.Println("defer_3")

	for i := 0; i < 3; i++ {
		// 在延迟调用后i=3，只输出3
		defer func() {
			fmt.Println(i)
		}()
	}

	open("book.doc")

	// panic/recover
	if throwPanic(p) {
		fmt.Println("Some error need to throw")
	}
}

func A(q int) int {
	fmt.Println("Func A")
	p := q
	return p
}

// 显示声明返回参数
func B(q ...int) (x, y int) {
	fmt.Println("Func B")
	fmt.Println(q)

	x = q[0]
	y = q[1]

	return
}

// 函数作为参数
func C(f func(int) int, k func(int)) {
	fmt.Println("函数做为参数")
	// 有返回值参数
	fmt.Println(f(88))
	// 无返回值函数
	k(88)
}

func Ok(k int) {
	fmt.Println("Ok", k)
}

// 闭包
func closure(x int) func(int) int {
	fmt.Println("Func closure")
	fmt.Println(&x)
	return func(y int) int {
		fmt.Println(&x)
		return x + y
	}
}

// defer
func open(file string) {
	fmt.Println("Open", file)
	defer close2()
	failure := failure()
	if failure {
		fmt.Println("Open failure...")
	}
}

func failure() bool {
	fmt.Println("Faiuer check...")
	res := false

	defer func(val bool) {
		res = val
	}(true)

	return res
}

func close2() {
	fmt.Println("Close...")
}

// panic/recover
func p() {
	fmt.Println("Invoke panic")
	x := "Painc in action"
	panic(x)
}

func throwPanic(p func()) (b bool) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("throwPanic", err)
			b = true
		}
	}()
	p()
	return
}