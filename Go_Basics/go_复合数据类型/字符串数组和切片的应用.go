package main

import "fmt"

/*字符串数组和切片的应用*/

//从字符串生成字节切片
func main_1() {
	s := "asdfaeafa"
	new_s := s[1:3]
	fmt.Println(new_s)
}

func main_2() {
	//修改字符串中的某一个字符
	/*
		Go 语言中的字符串是不可变的,需要吧字符串转换为字节数组，修改指定值后，在转会字符串
	*/

	s_1 := "sushenghong"
	new_s1 := []byte(s_1) //将字符串转换为字节数组
	new_s1[0] = 'H'       //修改指定值
	s_1 = string(new_s1)  //将字节数组转换为字符串
	fmt.Println(s_1)      //out Hushenghong
}
