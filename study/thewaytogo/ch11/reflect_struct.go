package main

import (
	"reflect"
	"fmt"
)

type UnknowType struct {
	s1, s2, s3 string
}

func (n UnknowType) ToString() string {
	return n.s1 + "-" + n.s2 + "-" + n.s3
}

var secret interface{} = UnknowType{"Ada", "Go", "Oberon"}

func main() {
	value := reflect.ValueOf(secret)
	//typ := value.Type()
	typ := reflect.TypeOf(secret)
	fmt.Println(typ)
	knd := value.Kind()
	fmt.Println(knd)

	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
	}

	result := value.Method(0).Call(nil)
	fmt.Println(result)
}