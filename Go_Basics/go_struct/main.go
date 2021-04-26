package main

/*
结构体名称大写也是对外可见的
结构体字段的可见性与json序列化

如果一个go语言包中定义的标识符是首字母大写的，那么就是对外可见的


*/

type Student struct {
	Id   int
	Name string
	Age  int
}

func main() {

}
