/*
  字符串
 */

package main

import "fmt"

func main() {
	s := "abc"
	println(s[0] == '\x61', s[1] == 'b', s[2] == '\x63')

	s = `a
	b\r\n\x00
	c
	`
	println(s, &s)

	s = "Hello, " +
		"World !"
	println(s)

	s1 := s[:5]
	s2 := s[7:]
	s3 := s[1:5]
	println(s1, s2, s3)

	fmt.Printf("%T\n", 'a')
	var c1, c2 rune = '\u6211', '们'
	println(c1 == '我', string(c2) == "\xe4\xbb\xac")

	//byte
	s = "abcd"
	bs := []byte(s)
	bs[1] = 'B'
	println(string(bs))

	//rune
	u := "电脑"
	rs := []rune(u)
	rs[1] = '话'
	println(string(u))

	s = "abc汉字"

	//byte
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c,", s[i])
	}

	fmt.Printf("\n")

	//rune
	for _, r := range s {
		fmt.Printf("%c,", r)
		fmt.Println(r)
	}
}
