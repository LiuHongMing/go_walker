package main

import (
	"fmt"
	"time"
)

func pump3() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck3(ch chan int) {
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
}

func main() {
	suck3(pump3())
	time.Sleep(1e9)
}
