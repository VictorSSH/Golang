package main

import (
	. "./add"   //表示该目录下所有的文件
	SUB "./sub" //给这个目录起个别名

	"fmt"
)

func main() {
	ret := SUB.Sub(2, 3)
	fmt.Println(ret)
	//fmt.Println("",add.Add(3, 4))
	Add(12, 12)
	fmt.Println(Show("北京", 12))
	fmt.Println(SUB.InfoShow("我这一生啊有太多坎坷"))

}
