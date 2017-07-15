package main

import "fmt"

// 声明struct
type person struct {
	Name    string
	Age     int

	// 内置struct
	Contact struct {
			Phone, City string
		}
}

type human struct {
	Sex int
}

type teacher struct {
	// 嵌套struct
	human

	Name string
	Age  int
}

type student struct {
	human
}

type A struct {
	B
}

type B struct {
	C
	Name string
}

type C struct {
	Name string
}

type TZ int

func main() {
	// struct 声明、赋值
	a := person{}
	a.Name = "go"
	a.Age = 10
	fmt.Println(a)
	a2 := person{Name: "go", Age: 0}
	fmt.Println(a2 == a)
	fmt.Println("--------------------")
	b := &person{Name: "gogo", Age: 20}
	fmt.Println(b)
	A1(b)
	B1(b)
	fmt.Println(b)
	fmt.Println("--------------------")
	// 匿名struct
	c := struct {
		Name string
		Age  int
	}{
		Name: "gogogo",
		Age: 30,
	}
	fmt.Println(c)
	fmt.Println("--------------------")
	d := person{Name: "jojo", Age: 34}
	d.Contact.Phone = "13691104030"
	d.Contact.City = "beijing"
	fmt.Println(d)
	fmt.Println("--------------------")
	x := teacher{Name: "teacher Li", Age: 40, human: human{Sex: 1}}
	x.Sex = 0
	fmt.Println(x)
	y := student{human{Sex:1}}
	fmt.Println(y)
	fmt.Println("--------------------")
	diff := A{B: B{Name: "b"}}
	diff.Name = "a"
	diff.B.Name = "bb"
	fmt.Println(diff.Name, diff.B.Name)
	// fmt.Println(diff.Name, diff.B.Name), 错误, 编译器无法识别是那个Name
	fmt.Println("--------------------")
	// 方法
	z := teacher{human: human{Sex: 12}}
	z.eat("fruit")
	fmt.Println(z)
	k := student{human{Sex:11}}
	k.eat("bread")
	fmt.Println(k)
	fmt.Println("--------------------")
	var tz TZ
	tz.Println() // method value
	(*TZ).Println(&tz) // method expression

	var tz2 TZ
	tz2.Increment(100)
	fmt.Println(tz2)

	// 分配内存, 声明指针
	var ps *person = new(person)
	ps.Name = "天行者"
	ps.Age = 88
	fmt.Println(ps)
	A1(ps)
	fmt.Println(ps)
}

func A1(a *person) {
	a.Name = "gogo888"
	fmt.Println(a)
}

func B1(b *person) {
	b.Age = 100
	fmt.Println(b)
}

// struct, 自定义类型添加方法
// (h *human) 表示receiver
func (h *human) eat(food string) {
	h.Sex = 101
	fmt.Println("Eat", food)
}

func (tz *TZ) Println() {
	fmt.Println("TZ")
}

func (tz *TZ) Increment(num int) {
	*tz += TZ(num)
}