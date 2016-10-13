package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	http.HandleFunc("/hello", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}