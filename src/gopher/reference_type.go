/*
  引用类型
 */

package main

func main() {
	a := []int{0, 0, 0}
	a[1] = 10

	b := make([]int, 3)
	b[1] = 10

	//c := new([]int)
	//c[1] = 10 // invalid opt
}
