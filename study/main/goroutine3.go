package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	go sendData3(ch)
	getData3(ch)
}

func sendData3(ch chan string) {
	ch <- "Beijing"
	ch <- "Shanghai"
	ch <- "Guangzhou"
	ch <- "Shenzhen"
	close(ch)
}

func getData3(ch chan string) {
	for {
		input, open := <-ch
		if !open {
			break
		}
		fmt.Printf("%s ", input)
	}
}