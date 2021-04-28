package main

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

func main1() {

	//匿名函数 声明后直接调用，
	func(name string) {
		fmt.Println("hell", name)
	}("上海")
}
