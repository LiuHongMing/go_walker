package main

import (
	"sort"
	"fmt"
)

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type IntArray []int

func (p IntArray) Len() int {
	return len(p)
}
func (p IntArray) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p IntArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

//func Sort(data Sorter) {
//	for pass := 1; pass < data.Len(); pass++ {
//		for i := 0; i < data.Len() - pass; i++ {
//			if (data.Less(i + 1, i)) {
//				data.Swap(i, i + 1)
//			}
//		}
//	}
//}
//
//func isSorted(data Sorter) bool {
//	n := data.Len()
//	for i := n; i > 0; i-- {
//		if (data.Less(i, i - 1)) {
//			return false
//		}
//	}
//	return true
//}

func main() {
	//实现sort.Interface接口才能直接使用sort.Sort函数进行排序
	//var data IntArray = []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	//sort.Sort(data)
	var data = []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	sort.Ints(data)
	fmt.Println(data)
}