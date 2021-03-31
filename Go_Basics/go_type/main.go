package main

import "fmt"

/*
		go中的数据类型

下面是 Go 支持的基本类型：

	bool
	数字类型
		int8, int16, int32, int64, int
		uint8, uint16, uint32, uint64, uint
		float32, float64
		complex64, complex128
		byte
		rune
	string

*/

/*golang中基本数据类型转换*/
func main() {
	var i int32 = 100
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	var n3 int64 = int64(i)
	fmt.Printf("i=%v\nn1=%v\nn2=%v\nn3=%v", i, n1, n2, n3)
	fmt.Printf("\nn1=%T\nn2=%T\nn3=%T", n1, n2, n3)

	/*转换时go 实际转换的是变量存储的值，而不是变量本身*/
	fmt.Printf("\ni type is %T\n", i)
	/*在转换中，比如将int64 转换为int8，此时编译时不会报错，只是转换的结果按照溢出处理，和我们期望值不一样*/
	var valu1 int64 = 3000000
	var valu2 int8 = int8(valu1)
	fmt.Printf("\nvalu1=%v\n", valu1)
	fmt.Printf("\nvalu2=%v\n", valu2)
	/*
		基本数据类型转换时需要注意一下三点
		1,go 中数据类型的转换可以是从表示范围小到表示范围大，也可以范围大到范围小
		2，被装欢的是变量的值，而不是变量本身，
		3，在转换中，比如将int64类型的数据转为int8类型的数据时，其编译不会报错，但是实际转换的结果并不是我们期望的结果
	*/
}
