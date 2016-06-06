package main

import "fmt"

func generate2() chan int {
	in := make(chan int)
	go func() {
		for i := 2; ; i++ {
			in <- i
		}
	}()
	return in
}

func filter2(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i % prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func sieve() chan int {
	out := make(chan int)
	go func() {
		ch := generate2()
		for {
			prime := <-ch
			ch = filter2(ch, prime)
			out <- prime
		}
	}()
	return out
}

func main() {
	primes := sieve()
	for {
		fmt.Print(<-primes, " ")
	}
}