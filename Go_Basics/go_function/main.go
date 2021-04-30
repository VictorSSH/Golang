package main

import (
	i "./init"
	u "./utils"
	f "fmt"
)

/**
	go 函数


--函数定义
函数名：由字母、数字、下划线组成。但函数名的第一个字母不能是数字。在同一个包内，函数名也称不能重名（包的概念详见后文）。
参数：参数由参数变量和参数变量的类型组成，多个参数之间使用,分隔。
返回值：返回值由返回值变量和其变量类型组成，也可以只写返回值的类型，多个返回值必须用()包裹，并用,分隔。
函数体：实现指定功能的代码块。

定义函数
//定义一个没有返回值的函数，并且包含两个int类型的变量
func func_01(x, y int) {
	f.Println(x, y)
}
定义两个变量的参数，并且返回值为int类型
func intSum(a, b int) (ret int) {
	ret = a + b
	return ret
}
				func showInf(name string, age int, salay int32) {
					f.Println("我叫：", name, "我今年:", age, "我工资是：", salay)

				}
函数调用
				func mainTest01() {
					//函数调用
					func_01(12, 43)
					intSum(2, 3)
					f.Println(intSum(2, 3))
					showInf("张三", 23, 34556)

				}
	函数调用机制




函数参数

函数的参数中如果相邻变量的类型相同，则可以省略类型
可变参数是指函数的参数数量不固定。Go语言中的可变参数通过在参数名后加...来标识。
注意：可变参数通常要作为函数的最后一个参数

*/

var (
	NumberOny  = 7
	NumberNy   = 40
	NumberNrti = 30
)

func main() {
	reslu := u.Cal(NumberOny, NumberNy, '+')
	reslu2 := u.Cal(NumberOny, NumberNrti, '+')
	f.Println(reslu, reslu2)

	reslu1 := u.GetSequence()
	f.Println(reslu1())

	/*go语句执行顺序*/
	n1 := 10
	u.Test01(n1)
	f.Println("main中的n1=", n1)

	/*计算两个数的和*/
	retu := u.Getsum(NumberNrti, NumberNy)
	f.Println("两数之和为= ", retu)
	f.Println("----------------------------")

	//sum, sub := u.Get_sum_sub(NumberNy, NumberOny)
	//f.Println("两数之和为:", sum)
	//f.Println("两数之差为：", sub)
	f.Println("----------------------------")
	sum1, _, fa := u.Get_sum_sub(NumberOny, NumberOny)
	f.Println(sum1, fa)
	f.Println("----------------------------")
	res := u.Billion(3)
	f.Println(res)
	f.Println(u.Billion(4))
	f.Println(u.Billion(6))
	f.Println(u.Billion(7))
	f.Println("----------------------------")
	female := u.Female(3)
	f.Println(female)
	f.Println("----------------------------")
	a := 10
	b := 20
	swap1, swap2 := u.Swap(&a, &b)
	f.Printf("swap1=%v\nswap2=%v\n", *swap1, *swap2)
	f.Println("----------------------------")
	f.Println(i.Id, i.Name, i.Info)
	id := i.Id
	name := i.Name
	info := i.Info
	id = 34
	name = "上海"
	info = "浦东新区"
	f.Println(id, name, info)
	f.Println("----------------------------")
	sub := i.FunallSum(1, 2)
	f.Println(sub)
	f.Println("----------闭包的实现-----------")

	num1 := 20
	i.Test1(&num1)
	f.Println("----------指针-----------")
	num2 := 30
	num3 := 10
	f.Println(i.Test2(&num3, &num2))

	/*
		&符号是指针取地址
		*符号是指针去地址值
	*/
	f.Println("----&符号取地址----")
	var i = 10
	f.Println(&i)
	f.Printf("i的；类型是: %T\n", i)
	f.Println("----*符号取地址值-----")
	var ptr *int
	f.Println("指针类型变量初始化值为 nil ", ptr)
	ptr = &i
	f.Printf("prt类型是%T\n", ptr)
	f.Printf("ptr通过*号来获取i变量的值: %d", *ptr)
}
