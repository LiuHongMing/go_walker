package main

import (
	"net/http"
	"log"
	"fmt"
	"strings"
)

func sayHelloWorld(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("form", r.Form)
	fmt.Println("request", r)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println("host", r.URL.Host)
	fmt.Println("path", r.URL.Path)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, " "))
	}
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	http.HandleFunc("/hello", sayHelloWorld)
	log.Fatal(http.ListenAndServe(":8080", nil))
}