package main

import (
	"time"
	"fmt"
)

func source(ch chan <- int) {
	for {
		ch <- 1
	}
}

func sink(ch <- chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func main() {
	ch := make(chan int) //bidirectional
	go source(ch)
	go sink(ch)
	time.Sleep(1e9)
}