/*
  数组
 */

package main

import "fmt"

func modify(array [5]int) {
	array[0] = 10
	fmt.Println("In modify, array values: ", array)
}

func main() {
	array := [5]int{5, 4, 3, 2, 1}
	modify(array)
	fmt.Println("In main, array values: ", array)
}
