package sub

import "fmt"

//在go语言中，同一层级目录，不允许出现多个相同的包名
func Sub(a, b int) int {
	fmt.Println("这是减法运算")
	return a - b
}
