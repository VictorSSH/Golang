package main

import (
	u "../utils"
	f "fmt"
)

/*go语言中函数也是一种数据类型
可以赋值给一个变量，则该变量就是一个函数类型变量，通过该变量可以对函数调用
*/

func getsum(n1, n2 *int) int {
	return *n1 + *n2
}
func main6() {

	a := getsum
	f.Println(a)
	f.Printf("a的类型是%T\ngetsunm类型是%T\n", a, getsum)

	res := a
	f.Printf("res的类型是%T\n,a的类型是%T\n", res, a)
}

func test(n1 int) {
	n1 = n1 + 1
	f.Println("test n1 = ", n1) //11
}
func test2(n1, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

var name string = "beijing"

func main() {
	//n1 := 10
	//test(n1)
	//fmt.Println("main n1 =", n1) //10
	//res1, res2 := test2(12, 5)
	//fmt.Println("sum =", res1, "sub= ", res2)
	//_, res2 = test2(23, 12)
	//fmt.Println(res2)
	/*根据键盘输出数字打印出相应金字塔*/
	//var n int
	//f.Println("请输入打印的层数：")
	//f.Scanln(&n) //Scanln中存储的是n变量的地址名称
	//u.PrintPaly(n)
	///*输出一个整数，表示要打印的乘法表*/
	//var num int
	//f.Println("请输入要打印的数：")
	//f.Scanln(&num)
	//u.PrintMuulti(num)

	u.Test02()

}
