package main

import (
	"net/http"
	"log"
	"fmt"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":11000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	os.Exit(0)
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}