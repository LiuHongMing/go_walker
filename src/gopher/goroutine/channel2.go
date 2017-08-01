/*
  单向channel
 */

package main

func Print(ch chan <- int) {
	ch <- 1
}

func main() {
	var ch1 chan int = make(chan int) // ch1正常的channel, 不是单向
	var ch2 chan <- float64 // ch2是单向channel，只用于写float64数据
	var ch3 <- chan int     // ch3是单向channel，只用于读int数据

	ch4 := make(chan <- int) // ch4是单向channel，只用于写int数据

	println(ch1, ch2, ch3, ch4)

	go Print(ch4)

	value, ok := <-ch4 // invalid operation: <-ch4 (receive from send-only type chan<- int)
	if ok {
		println("ch4 output value is", value)
	} else {
		println("ch4 is a input direct channel")
	}
}