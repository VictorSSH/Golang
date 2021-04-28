package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}
func colc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func a() func() {
	return func() {
		fmt.Println("hello")
	}
}

func b() func() {
	name := "北京"
	return func() {
		fmt.Println("hekll", name)
	}
}

func sgh(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(y int) int {
		base += y
		return base
	}

	return sub, add

}

func main() {

	//	p1 := colc(2,4,add)
	//	fmt.Println(p1)
	//	p2 := colc(4,2,sub)
	//	fmt.Println(p2)
	// p1 := a()
	// p1()
	//
	// p2 := b()
	// p2()

	// ls,hj := sgh(1)
	//a :=ls(2)
	// b:=hj(4)
	//
	// fmt.Println(a)
	// fmt.Println(b)
}
