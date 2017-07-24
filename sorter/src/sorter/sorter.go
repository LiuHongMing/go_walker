package main

import "flag"

func main() {
	var infile *string = flag.String("i", "infile", "File contains values for sorting")
}