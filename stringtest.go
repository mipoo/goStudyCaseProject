package main

import "fmt"

// 修改字符串的错误示例
func main() {
	x := "你好"
	xBytes := []byte(x)
	xBytes[0] = 'T'	// 注意此时的 T 是 rune 类型
	x = string(xBytes)
	fmt.Println(x)	// Text

	x2 := "你好"
	xRunes := []rune(x2)
	xRunes[0] = 'T'
	x2 = string(xRunes)
	fmt.Println(x2)	// 我ext
}

