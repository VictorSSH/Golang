package demo01

import (
	f "fmt"
)

func Test() {
	var arr1 [5]int
	for i := 0; i < len(arr1); i++ {
		arr1[0] = i * 2
	}

	for i := 0; i < len(arr1); i++ {
		f.Printf("arr1 at index %d is %d\n", i, arr1[i])
	}

}
func Test2() {
	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		f.Printf("array item", i, "is", a[i])
	}
}
func Test3(val []int) {
	for v := range val {
		f.Println(v)
	}

}

func intTestarray01() {

}
