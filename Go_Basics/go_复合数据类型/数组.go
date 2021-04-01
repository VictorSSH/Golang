package main

import "fmt"

/*
	go 中的数组
	概念：
		数组是一组具有相同唯一类型的一组以编号且长度固定的数据序列
	数组元素可以通过 索引（位置）来读取（或者修改），索引从 0 开始，第一个元素索引为 0，第二个索引为 1，以此类推。
	（数组以 0 开始在所有类 C 语言中是相似的）。元素的数目，也称为长度或者数组大小必须是固定的并且在声明该数组时就给出
	（编译时需要知道数组长度以便分配内存）；数组长度最大为 2Gb。

	声明格式
	var identifier [len]type
	例如
	var arr1 [6]int
	说明
		每个元素是一个整型值，当声明数组时所有的元素都会被自动初始化为默认值 0。
		arr1 的长度是 5，索引范围从 0 到 len(arr1)-1。
		第一个元素是 arr1[0]，第三个元素是 arr1[2]；总体来说索引 i 代表的元素是 arr1[i]，最后一个元素是 arr1[len(arr1)-1]。
		对索引项为 i 的数组元素赋值可以这么操作：arr[i] = value，所以数组是 可变的。
*/

/**
由于索引的存在，遍历数组的方法自然就是使用 for 结构:
通过 for 初始化数组项
通过 for 打印数组元素
通过 for 依次处理元素
*/

/*定义*/

func main() {
	/*声明格式
	var identifier [len]type*/

	//声明指定长度的数组
	arrname1 := [5]int{1, 3, 544, 565, 787}
	arrname2 := [3]string{"北京", "上海", "兰州"}

	//声明任意长度数组

	arrname3 := [...]int{32, 43, 546, 76, 87, 8545}
	arrname4 := [...]string{"shanghai", "beijing", "lanzhou", "yunna", "hebei"}

	//输出全部数组元素
	fmt.Println(arrname1, arrname2, arrname3, arrname4)

	//输出指定下标索引位置元素
	fmt.Printf("数组元素值:%s\n", arrname4[1])

	//循环方式进行数组元素遍历

	fmt.Println("----------------------循环方式输出遍历所有数组元素-----------------")
	for i := 0; i < len(arrname4); i++ {
		fmt.Println(i, arrname4[i])
	} /*
		输出


		0 shanghai
		1 beijing
		2 lanzhou
		3 yunna
		4 hebei

	*/
}

/*循环方式进行数组初始化和遍历元素值*/
func main1() {

	var arr1 [5]int
	//通过for 循环进行arr[] 数组发赋值
	for i := 0; i < len(arr1); i++ {
		//循环一次，i乘以2
		arr1[i] = i * 2
	}
	//通过for循环来进行对数组输出
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("数组索引:%d is %d\n", i, arr1[i])
	} /*
		输出：
		数组索引:0 is 0
		数组索引:1 is 2
		数组索引:2 is 4
		数组索引:3 is 6
		数组索引:4 is 8
	*/

	/*也可以使用for renge方式实现对数组的操作*/

	for key, valu := range arr1 {
		fmt.Printf("数组索引：%d is 数组值: %d\n", key, valu)
	} /*
		输出：
		数组索引：0 is 数组值: 0
		数组索引：1 is 数组值: 2
		数组索引：2 is 数组值: 4
		数组索引：3 is 数组值: 6
		数组索引：4 is 数组值: 8
	*/

	a := [...]string{"a", "b", "c"}

	for i := 0; i < len(a); i++ {
		fmt.Println("数组", i, "is", a[i])
	} /*
		输出
		数组 0 is a
		数组 1 is b
		数组 2 is c
	*/

}

func main2() {
	/*同样类型的数组是可以相互赋值的，不同类型的不行，会编译错误。那么什么是同样类型的数组呢？Go语言规定，必须是长度一样，并且每个元素的类型也一样的数组，才是同样类型的数组。*/

	arr1 := [2]int{3, 6}
	var arr2 = arr1
	fmt.Println(arr2)

	var arr3 [2]int = arr1
	fmt.Println(arr3)

	/*一个指针变量可以指向任何一个值的内存地址 它指向那个值的内存地址，在 32 位机器上占用 4 个字节，在 64 位机器上占用 8 个字节，并且与它所指向的值的大小无关。
	当然，可以声明指针指向任何类型的值来表明它的原始性或结构性；你可以在指针类型前面加上 * 号（前缀）来获取指针所指向的内容，这里的 * 号是一个类型更改器。
	使用一个指针引用一个值被称为间接引用。*/
	/*

		Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。
		在指针类型前面加上 * 号（前缀）来获取指针所指向的内容，这里的 * 号是一个类型更改器。使用一个指针引用一个值被称为间接引用

	*/

	prt := 5
	fmt.Printf("An integer: %d,its location in memory: %p\n", prt, &prt) //（指针的格式化标识符为 %p）
	var intp *int
	intp = &prt
	fmt.Println(intp, *intp)
	fmt.Printf("类型为: %T", intp)

}
