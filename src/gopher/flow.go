/*
  流程控制
 */

package main

import "fmt"

func main() {
	myfunc()
}

// goto
func myfunc() {
	i := 0
	HERE:
	fmt.Println(i)
	i++
	if i < 10 {
		goto HERE
	}
}