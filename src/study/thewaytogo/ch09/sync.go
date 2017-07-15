package main

import (
	"sync"
	"fmt"
)

type Info struct {
	mu sync.Mutex
	Str string
}

func main() {
	var info Info = Info{Str: "Sync Mutex"}
	fmt.Println(info.Str)
	Update(&info)
	fmt.Println(info.Str)
}

func Update(info *Info) {
	info.mu.Lock()
	info.Str = "Update by sync mutex"
	info.mu.Unlock()
}
