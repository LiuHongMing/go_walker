/*
  数组分片
 */

package main

import "fmt"

func main() {
	var myArray = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Elements of myarray:")
	for _, v := range myArray {
		fmt.Print(v, " ")
	}

	fmt.Println()

	var mySlice = myArray[:5]
	fmt.Println("Elements of myslice:")
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}

	mySlice1 := make([]int, 5)
	mySlice2 := make([]int, 5, 10)
	mySlice3 := []int{1, 2, 3, 4, 5}
	println(mySlice1, mySlice2, mySlice3)
}
