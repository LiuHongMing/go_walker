package main

import (
	"fmt"
	"time"
	"sort"
)

func compute(ch chan int) {
	ch <- sumCompute()
}

func sumCompute() (sum int) {
	for i := 1; i < 10; i++ {
		sum += i
	}
	return
}

func doSomething() {
	time.Sleep(5 * 1e9)
}

func main() {
	ch := make(chan int)
	//go compute(ch)
	//doSomething()
	//result := <- ch
	//fmt.Println(result)

	go func() {
		ch <- 1
	}()
	doSomething()
	fmt.Println(<-ch)

}