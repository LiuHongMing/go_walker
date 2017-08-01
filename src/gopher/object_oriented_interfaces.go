/*
  接口：
      实现接口
      定义接口
      类型断言
 */

package main

import (
	"sort"
	"fmt"
)

// 实现了sort.Interface接口方法，Len、Less、Swap
type Sequence []int

// 被sort.IntSlice(seq).Sort()代替
//func (seq Sequence) Len() int {
//	return len(seq)
//}
//
//func (seq Sequence) Less(i, j int) bool {
//	return seq[i] < seq[j]
//}
//
//func (seq Sequence) Swap(i, j int) {
//	seq[i], seq[j] = seq[j], seq[i]
//}

func (seq Sequence) String() string {
	sort.IntSlice(seq).Sort()
	// 被sort.IntSlice(seq).Sort()代替
	// sort.Sort(seq)

	// 被fmt.Sprint([]int(seq))代替
	//str := "["
	//for i, elem := range seq {
	//	if i > 0 {
	//		str += " "
	//	}
	//	str += fmt.Sprint(elem)
	//}
	//return str + "]"
	return fmt.Sprint([]int(seq))
}

// 定义, Sequence实现了Stringer接口的String()方法
type Stringer interface {
	String() string
}

func main() {
	var seq Sequence = Sequence([]int{1, 3, 2})
	println(seq.String())

	// 接口转换，做为任意类型
	var val interface{} = seq
	println("val:", toString(val))

	// 类型、接口查找
	v, ok := val.(Stringer);
	if ok {
		println("val.(Stringer):", v.String())
	}

	var val2 interface{} = 10
	println("val2:", toString(val2))
}

func toString(val interface{}) string {
	// 类型断言
	switch str := val.(type) {
	case string:
		return str
	case Stringer:
		return str.String()
	default:
		return "Neither string nor Stringer"
	}
}