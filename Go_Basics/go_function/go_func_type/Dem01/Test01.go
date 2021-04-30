package Dem01

/*
	go函数类型变量

	 可以使用type 关键字来定义一个函数类型，
	type calculation func(int,int)int

上面语句定义了一个calculation类型，它是一种函数类型，这种函数接收两个int类型的参数并且返回一个int类型的返回值。

简单来说，凡是满足这个条件的函数都是calculation类型的函数，例如下面的add和sub是calculation类型
*/

type My_func func(int, int) int

func Add(a, b int) int {
	return a + b
}

func Sub(a1, a2 int) int {
	return a1 - a2
}

/*
	这个intSeq,函数返回另一个在intSeq函数体内定义的匿名函数
	这个返回的函数使用闭包的方式_隐藏_ 变量 i
*/

func IntSeq() func() int {
	i := 2
	return func() int {
		i++
		return i
	}
}

func Test1() {

}
