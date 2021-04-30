package main

import (
	t "./Dem01"
	f "fmt"
)

func test1() {
	defer func() {
		err := recover()
		if err != nil {
			f.Println("err=", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	f.Println("res=", res)
}

func main() {
	var newmy_func t.My_func
	newmy_func = t.Add

	sub := t.Sub(5, 2)
	f.Println(sub)
	f.Printf("newmy_func类型是: %T\n", newmy_func)
	//调用该函数
	f.Println(newmy_func(2, 3))

	nextInt := t.IntSeq()
	f.Println(nextInt)
	test1()
	f.Println("代码继续执行")
}
