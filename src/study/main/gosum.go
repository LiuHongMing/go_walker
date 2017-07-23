package main

import "fmt"

func sum(x, y int, c chan int) {
	c <- x + y
}

func main() {
	c := make(chan int)
	go sum(6, 8, c)
	fmt.Println(<-c)
}