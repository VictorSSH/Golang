package main

import "fmt"

/*
	go 中自定义类型

	在Go语言中有一些基本的数据类型，如string、整型、浮点型、布尔等数据类型， Go语言中可以使用type关键字来定义自定义类型。
			//将MyInt定义为int类型
			type MyInt int

*/
/*自定义类型*/
type myint int

/*类型别名*/
type newmyint = myint

func main() {
	var i myint
	fmt.Printf("type:%T values:%v\n", i, i)
	var name newmyint
	fmt.Printf("type:%T values:%v\n", name, name)
}
