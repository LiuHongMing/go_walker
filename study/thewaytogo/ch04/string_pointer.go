package main

import "fmt"

func main() {
	s := "Good bye"
	var p *string = &s
	*p = "ciao"
	fmt.Println("Here is the pointer p", p)
	fmt.Println("Here is the string *p", *p)
	fmt.Println("Here is the string s", s)
}