package vict

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

func Str_u() {

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
func main2() {

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
	/*只是创建了一个emp结构体类型，并没有定义结构体类型*/
	emp01 := struct {
		a1 int
		a2 struct{}
		a3 string
		a4 bool
		a5 float64
		a6 byte
		a7 rune
	}{}
	fmt.Println(emp01) /*

		输出结果为：

			{0 {}  false 0 0 0}
	*/
}

/*访问结构体中字段值*/

type name_01 struct {
	name string
	age  int
	addr string
}

func main3() {

	/*第一种方式初始化值*/
	n1 := name_01{"王五", 12, "上海市"}
	fmt.Printf("名称%s\n", n1.name)
	fmt.Printf("年龄%d\n", n1.age)
	fmt.Printf("地址%s\n", n1.addr) /*
		输出
		名称王五
		年龄12
		地址上海市
	*/

	/*第二种方式初始化值*/

	var n2 name_01
	n2.name = "赵四"
	n2.age = 12
	n2.addr = "上海市"
	fmt.Println(n2)
	/*分开获取每个字段中值*/
	fmt.Printf("姓名:%s\n年龄%d\n地址%s\n", n2.name, n2.age, n2.addr) /*
		 输出：
		 	{赵四 12 上海市}

			姓名:赵四
			年龄12
			地址上海市
	*/
}

/*结构体指针*/

type info_01 struct {
	val1 string
	val2 int
	val4 float64
}

func main4() {
	info := &info_01{"北京", 32, 32}
	fmt.Println((*info).val1, (*info).val2, (*info).val4) /*
		结果
			北京 32 32
	*/
}
