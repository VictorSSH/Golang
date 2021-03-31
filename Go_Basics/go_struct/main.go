package main

import "fmt"

/*
结构体是用户定义的类型，表示若干个字段（Field）的集合。
有时应该把数据整合在一起，而不是让这些数据没有联系。这种情况下可以使用结构体。
*/

/*声明结构体*/
type Employee struct {
	firstName string
	lastName  string
	age       int16
}

//上面定义了一个Employee 类型的结构体，具有三个字段属性，其中firstName，lastName字段的类型都是string类型，可以简写为一下格式
type Employee1 struct {
	firstName, lastName string
	age, salary         int16
}

/*匿名结构体*/
/*
	在声明结构类型是，可以不指定该结构类型称为匿名结构体
*/
var Employee2 struct {
	firstName, lastName string
	age, salary         int16
}

func main1() {

	emp1 := Employee{
		firstName: "张三",
		lastName:  "it",
		age:       12,
	}
	fmt.Println(emp1)
	fmt.Printf("名称：%s\n职位：%s\n年龄：%d\n", emp1.firstName, emp1.lastName, emp1.age)

	emp2 := Employee{"张三", "文员", 22} //这种方式赋值，需要改结构字段顺序一致，才行
	fmt.Println("emp1:", emp1)
	fmt.Println("emp2", emp2)
}

/*创建匿名结构体*/

func main() {

	emp := struct {
		firstName, lastName string
		age, salary         int
		//注意在创建匿名结构时，字段类型后面不需要加分号，而在非匿名结构体时，需要在字段类型后面加分号
	}{
		firstName: "李四",
		lastName:  "司机",
		age:       23,
		salary:    34000,
	}
	fmt.Println(emp)
}
