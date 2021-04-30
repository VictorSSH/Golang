package utils

import (
	"fmt"
)

func f1() {
	fmt.Println("hello 上海")
}
func f2(name1, name2 string) {
	fmt.Println(name1)
	fmt.Println(name2)
}

//可变参数,可以使用for循环变量
func f3(names ...string) {
	fmt.Println(names)
}

/*函数多返回值*/
/*Go语言支持函数方法的多值返回，也就说我们定义的函数方法可以返回多个值，比如标准库里的很多方法，都是返回两个值，第一个是函数需要返回的值，
第二个是出错时返回的错误信息，这种的好处，我们的出错异常信息再也不用像Java一样使用一个Exception这么重的方式表示了，非常简洁*/

func hardEight() {

	//匿名函数 声明后直接调用，
	func(name string) {
		fmt.Println("hell", name)
	}("上海")
}

/*
	使用函数递归方式，求出斐波那契数列，
	给出一个整数n，求出它的斐波那契数列是多少
*/
func Billion(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return Billion(n-1) + Billion(n-2)
	}
}

/*
	求函数值已知 f(1)=3; f(n)= 2 *f(n-1)+1,请使用递归方式求出f(n)的值
*/
func Female(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*Female(n-1) + 1
	}

}

/*
	编写一个函数，来交换两个变量的值
*/

func Swap(n1 *int, n2 *int) (*int, *int) {
	t := *n1
	*n1 = *n2
	*n2 = t
	return n1, n2
}

func Test02() {
	for a := 0; a < 10; a++ {
		fmt.Println(a)
	}
}
