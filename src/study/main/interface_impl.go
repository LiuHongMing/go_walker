package main

import "fmt"

type S struct {
	i int
}

func (s *S) Get() int {
	return s.i
}

func (s *S) Put(v int) {
	s.i = v
}

type newS S

type Simp struct {
	newS
}

type I interface {
	Get() int
	Put(int)
}

func f(p I) {
	fmt.Println(p.Get())
	p.Put(1)
}

func main() {
	var s S = S{10}
	fmt.Println(s.Get())
	s.Put(20)
	fmt.Println(s.i)

	var ns newS = newS{50}
	fmt.Println(ns)

	var simp Simp = Simp{newS: newS{i: 100}}
	simp.i = 200
	fmt.Println(simp)

	var i I = &S{i: 96}
	fmt.Println(i)
	i.Put(296)
	fmt.Println(i)
}

type TK int

func (tk *TK) Show() {
	*tk = 2
}