package main

import (
	"fmt"
	"errors"
)

var m, d = 4, 0

func main() {
	// if
	PrintIf()
	a := [...]int{1, 2, 3, 4, 5}
	fmt.Println(a)

	// for
	// reverse a
	for i, j := 0, len(a) - 1; i < j; i, j = i + 1, j - 1 {
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println(a)

	list := []string{"a", "b", "c", "d", "e"}
	for i, _ := range list {
		list[i] += "1"
	}
	fmt.Println(list)

	for i, c := range "go" {
		fmt.Println(i, c)
	}

	// switch
	var grade string = "B"
	var marks int = 80

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch grade {
	case "A":
		fmt.Print("优秀\n")
	case "B", "c":
		fmt.Print("良好\n")
	case "D":
		fmt.Print("及格\n")
	default:
		fmt.Print("不及格\n")
	}
	fmt.Printf("你的等级：%s\n", grade)

	// type match
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Println("x的类型：", i)
	case int:
		fmt.Println("x的类型：int")
	case float64:
		fmt.Println("x的类型：float64")
	case func(int):
		fmt.Println("x的类型：func(int)")
	default:
		fmt.Println("Unknow...")
	}
}

func PrintIf() {
	v, err := Division(m, d)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
}

func Division(m int, d int) (int, error) {
	if d == 0 {
		return 0, errors.New("除数为0")
	}
	return m / d, nil
}
