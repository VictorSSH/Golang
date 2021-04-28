package main

import "fmt"

/*go语言中函数也是一种数据类型
可以赋值给一个变量，则该变量就是一个函数类型变量，通过该变量可以对函数调用
*/

func getsum(n1, n2 *int) int {
	return *n1 + *n2
}
func main6() {

	a := getsum
	fmt.Println(a)
	fmt.Printf("a的类型是%T\ngetsunm类型是%T\n", a, getsum)

	res := a
	fmt.Printf("res的类型是%T\n,a的类型是%T\n", res, a)
}

func test(n1 int) {
	n1 = n1 + 1
	fmt.Println("test n1 = ", n1) //11
}
func test2(n1, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}
func main() {
	n1 := 10
	test(n1)
	fmt.Println("main n1 =", n1) //10
	res1, res2 := test2(12, 5)
	fmt.Println("sum =", res1, "sub= ", res2)
	_, res2 = test2(23, 12)
	fmt.Println(res2)

}
