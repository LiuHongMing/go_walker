package mymath

import "errors"

// 小写字母开头的函数只在本包内可见
// 大写字母开头的函数才能被其他包使用
func Add(a, b int) (ret int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("Should be non-negative numbers!")
		return
	}
	return a + b, nil
}