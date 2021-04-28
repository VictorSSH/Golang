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
