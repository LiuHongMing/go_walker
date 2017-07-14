package main

import "fmt"

func main() {
	// array
	var balance = []int{1000, 20, 5, 10, 60}
	fmt.Println(balance)

	var avg float32
	avg = getAvg(balance, 5)

	fmt.Println("平均值：", avg)

	// slice
	a := [5]int{1, 2, 3, 4, 5}
	s1 := a[1:5]
	s2 := a[:]
	s3 := a[:4]
	s4 := s2[1:3]
	fmt.Println(s1, s2, s3, s4)

	// slice, 内置函数append
	sc1 := []int{0, 0}
	sc2 := append(sc1, 1, 2)
	sc3 := append(sc2, 3, 4, 5)
	sc4 := append(sc3, sc1...)
	fmt.Println(sc1, sc2, sc3, sc4)

	// slice, 内置函数copy
	var aa = [...]int{6, 7, 8, 9, 0, 4, 1, 3}
	var s = make([]int, 6, 10)
	n1 := copy(s, aa[0:])
	fmt.Println(n1, s)
	n2 := copy(s, aa[5:])
	fmt.Println(n2, s)

	// map
	var m = make(map[string]int) // 只声明map
	fmt.Println(m)

	mondays := map[string]int{
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
		"Oct": 31, "Nov": 30, "Dec": 31,
	}
	fmt.Println(mondays["Dec"])
	year := 0
	for _, days := range mondays {
		year += days
	}
	fmt.Println("Numbers of days in a year", year)
	mondays["Undecim"] = 30
	mondays["Feb"] = 29
	fmt.Println(mondays)

	v, ok := mondays["Jan"]
	fmt.Println(v, ok)

	delete(mondays, "Jan")
	fmt.Println(mondays)

	ma := map[string][]string{
		"x" : {"1", "2", "3"},
	}
	fmt.Println(ma)
}

func getAvg(arr []int, size int) float32 {
	var i, sum int
	var avg float32
	for i = 0; i < size; i++ {
		sum += arr[i]
	}
	avg = float32(sum / size)
	return avg
}