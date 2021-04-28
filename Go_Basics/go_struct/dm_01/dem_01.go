package main

import "fmt"

type class struct {
	Name string
	Age  int
	Add  string
}

func main() {

	//声明方式一
	var c1 class
	c1.Name = "小明"
	c1.Add = "北京市海淀区"
	c1.Age = 12
	fmt.Println(c1)

	//声明方式二

	var c2 *class = &class{
		Name: "小张",
		Age:  23,
		Add:  "北京市海淀区",
	}
	fmt.Println(c2)

}
