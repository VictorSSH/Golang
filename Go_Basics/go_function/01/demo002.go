package main

import (
	fm "fmt"
)

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func cal(n1, n2 int, operator byte) int {
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
		fm.Println("操作符错误..")
	}
	return res
}

func main() {

	resul := cal(5, 2, '-')
	resul1 := cal(5, 2, '+')
	resul2 := cal(5, 2, '*')
	fm.Println(resul, resul1, resul2)

	/*
		nexitNumber := getSequence()
		fm.Println(nexitNumber())
	*/
}
