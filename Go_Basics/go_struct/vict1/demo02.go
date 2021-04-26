package vict1

import (
	"fmt"
)

/*结构体实例化
只有当结构体实例化时，才会真正地分配内存。也就是必须实例化后才能使用结构体的字段。
结构体本身也是一种类型，我们可以像声明内置类型一样使用var关键字声明结构体类型。
var 结构体实例 结构体类型
*/

type person struct {
	name string
	age  int
	city string
	info struct{}
}

func main1() {
	var p1 = person{name: "北京", age: 12, city: "上海"}
	fmt.Println(p1)
}

//func main()  {
//	/*创建指针类型的结构体*/
//	var p2 = new(person)
//	fmt.Printf("%T\n",p2)
//	fmt.Printf("%#v\n",p2)
//
//
//}
