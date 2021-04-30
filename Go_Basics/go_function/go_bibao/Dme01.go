package go_bibao

import "fmt"

/*
	闭包的实现
*/

func Aupper() func(int) int {
	var n int = 10
	var str = "hello"
	return func(i int) int {
		n = n + i
		str += string(36)
		fmt.Println("str=", str)
		return n
	}
}
