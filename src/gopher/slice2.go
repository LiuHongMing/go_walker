/*
  数组分片2
 */

package main

import "fmt"

func main() {
	mySlice := make([]int, 5, 10)

	fmt.Println("len(mySlice):", len(mySlice))
	fmt.Println("cap(mySlice):", cap(mySlice))
	fmt.Printf("%#v\n", mySlice)

	//动态增减元素
	mySlice = append(mySlice, 1, 2, 3)
	mySlice2 := []int{8, 9, 10}
	mySlice = append(mySlice, mySlice2...)

	fmt.Printf("%#v\n", mySlice)

	//基于数组切片创建数组切片
	oldSlice := []int{1, 2, 3, 4, 5}
	newSlice := oldSlice[:3]
	fmt.Printf("%#v\n", newSlice)

	//内容复制
	mySlice3 := []int{1, 2, 3, 4, 5}
	mySlice4 := []int{5, 4, 3}

	copy(mySlice4, mySlice3)
	copy(mySlice3, mySlice4)

	fmt.Printf("%#v\n", mySlice3)
	fmt.Printf("%#v\n", mySlice4)
}
