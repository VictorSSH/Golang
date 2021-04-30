package init

import (
	f "fmt"
)

//定义全局变量

var Id int
var Name string
var Info string
var n3 int
var n4 int

//通常可以在init 函数中完成初始化操作
func init() {
	f.Printf("初始化init 函数")
	Id = 12
	Name = "北京市海地区"
	Info = "五道口广场"
}

//定义全局匿名函数

var (
	//求和
	FunallSum = func(n1, n2 int) int {
		return n1 + n2
	}
	//差值
	FunallSub = func(n3, n4 int) int {
		return n3 - n4
	}
	//乘法
	Funallmax = func(a1, a2 int) int {
		return a1 * a2
	}
)

func Test1(n1 *int) {
	*n1 = *n1 + 10
	f.Println("Test1 的 n1 ", *n1)

}
func Test2(n2, n3 *int) int {
	*n2 = 10
	*n3 = 10
	n4 = *n3 + *n2
	return n4
}
