package vict

import "fmt"

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

//定义方法

func (p Person) Dream() {
	fmt.Println("%s谁的梦想是学号go语言\n", p.name)
}
func main() {

	p1 := newPerson("sushenghong", "上海", 23)
	p1.Dream()
	fmt.Println(p1)
}
