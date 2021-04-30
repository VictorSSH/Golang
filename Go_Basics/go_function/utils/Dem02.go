package utils

import (
	f "fmt"
)

func GetSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

/*
	该函数体返回两个参数，用来计算
*/

func Cal(n1, n2 int, operator byte) int {
	var res int
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		f.Println("操作符错误..")
	}
	return res
}

func Test01(n1 int) {
	n1 = n1 + 1
	f.Println("函数Test01中的n1= ", n1)
}

/*计算两个数的和，并且返回值类型为int*/
func Getsum(value1, value2 int) int {
	return value1 + value2

}

/*函数计算两个数的和以及两个数的差，返回值类型为int类型*/
func Get_sum_sub(val1, val2 int) (int, int, int) {
	sum := val1 + val2
	sub := val1 - val2
	if sum <= 10 {
		f.Println("sum输入数太小")
		if sub <= 5 {
			f.Println("sub 数字太小")
		}
	}
	dha := Getsum(2, 2)
	return sum, sub, dha
}

/*使用函数输出对应的金字塔*/
func PrintPaly(totalevel int) {
	//var totalevel int = 20
	//i表示层数
	for i := 1; i <= totalevel; i++ {
		//输出空格
		for k := 1; k <= totalevel-i; k++ {
			f.Print(" ")
		}
		//j标识每层打印多少*
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == totalevel {
				f.Print("*")
			} else {
				f.Print(" ")
			}
		}
		f.Println()
	}
}

/*使用函数实现九九乘法表*/
func PrintMuulti(num int) {
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			f.Printf("%v * %v = %v \t", j, i, j*i)
		}
		f.Println()
	}
}
