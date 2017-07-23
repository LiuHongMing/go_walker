/**
  变量
 */

package main

//声明
var (
	v1 int
	v2 string
	v3 [10]int //数组
	v4 []int   //数组分片
)

var v5 struct {
	f int
}
var v6 *int           //指针
var v7 map[string]int //map.key为string, value为int
var v8 func(a int) int

func main() {
	//匿名变量
	_, _, nickName := GetName()
	println(nickName)

	//初始化
	var k1 int = 10
	var k2 = 10 //编译器自动推导类型
	k3 := 10    //编译器自动推导类型
	println(k1, k2, k3)

	//赋值
	v1 = 10
	println(v1)
}

func GetName() (firstName, lastName, nickName string) {
	return "liu", "hongming", "jason"
}
