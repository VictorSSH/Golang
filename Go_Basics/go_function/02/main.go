package main

import (
	"fmt"
	"os"
)

//定义图书属性
type book struct {
	title   string
	author  string
	price   float32
	publish bool
}

//定义菜单显示

func showInfo() {
	fmt.Println("欢迎登陆图书管理系统")
	fmt.Println("1, 添加书籍")
	fmt.Println("2, 修改书籍")
	fmt.Println("3, 显示书籍")
	fmt.Println("4, 退出")
}

//添加书籍函数
func addBooks() {}

func main2() {
	//使用for循环进行循环输出
	for {
		showInfo()
		//等待用户输出
		var option int
		fmt.Scanln(&option)
		switch option {
		case 1:
			addBooks()
		case 2:
			//updateBooks()updateBooks
		case 3:
			//showBooks()
		case 4:
			os.Exit(0)

		}
	}

}
