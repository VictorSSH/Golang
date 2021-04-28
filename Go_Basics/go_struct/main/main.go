package main

import (
	"fmt"
)

/*Person结构体*/
type Person struct {
	name string
	city string
	age  int8
}

/*构造函数
newPerson 是一个Person类型的构造函数
*/
func newPerson(name, city string, age int8) *Person {
	return &Person{
		name: name,
		city: city,
		age:  age,
	}
}
func newPerson_1(name string, age int8) *Person {
	return &Person{name: name, age: age}
}

//定义方法

func (p Person) Dream() {
	fmt.Printf("%s的梦想是学号go语言\n", p.name)
}

func (ectp Person) ect() {
	fmt.Println("学习也得要吃饭啊!", ectp.name, ectp.age)
}
func main12() {

	p1 := newPerson("sushenghong", "上海", 23)
	p1.Dream()

	p2 := newPerson_1("速升鸿", 23)
	fmt.Println((*p2).ect)

}
