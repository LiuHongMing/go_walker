package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 50)
	go func() {
		time.Sleep(5 * 1e9)
		x := <-ch
		fmt.Println("received", x)
	}()

	fmt.Println("sending", 10)
	ch <- 10
	fmt.Println("sent", 10)
}
