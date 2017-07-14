package main

import "fmt"

type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

type Append interface {
	Append(int)
}

func CountInto(a Append, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEngouh(l Lener) bool {
	return l.Len() * 10 > 42
}

func main() {
	var lst List
	//CountInto(lst, 1, 10) // 无效调用用,CountInto的recevier是List指针类型
	if LongEngouh(lst) {
		fmt.Println("- lst is long engouh")
	}

	plst := new(List)
	CountInto(plst, 1, 10)
	if LongEngouh(plst) {
		fmt.Println("- plst is long engouh")
	}
}
