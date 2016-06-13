package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig string = "666"
	var an int
	var newS string

	fmt.Println("The computer bitesize is ", strconv.IntSize)

	an, _ = strconv.Atoi(orig)
	fmt.Println(an)
	an += 5
	fmt.Println(an)
	newS = strconv.Itoa(an)
	fmt.Println(newS)
}
