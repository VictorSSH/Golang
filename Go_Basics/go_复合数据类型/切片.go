package main

import (
	"fmt"
)

/*
	概念
		切片是由数组实现的。它的底层是数组，可以理解为对底层数组的抽象，
		切片对象非常小，是因为它只有三个字段的数据结构，一个是指向底层数组的指针，一个切片的长度，一个切片的容量

	切片（slice）是对数组一个连续片段的引用（该数组我们称之为相关数组，通常是匿名的），所以切片是一个引用类型
	切片是可索引的，并且可以由 len() 函数获取长度
	切片提供了计算容量的函数 cap() 可以测量切片最长可以达到多少：它等于切片的长度 + 数组除切片之外的长度。如果 s 是一个切片，cap(s)
		就是从 s[0] 到数组末尾的数组长度。
	切片的长度永远不会超过它的容量，所以对于 切片 s 来说该不等式永远成立：0 <= len(s) <= cap(s)。
	优点 因为切片是引用，所以它们不需要使用额外的内存并且比使用数组更有效率，所以在 Go 代码中 切片比数组更常用

		声明切片
  			var identifier []type
	一个切片在未初始化之前默认为 nil，长度为 0
	切片的初始化格式是：var slice1 []type = arr1[start:end]

	创建切片 使用make方式创建切片

   slice0 :=	make([]int,3)	//使用make创建切片时需要传入一个参数，指定切片的长度，

    make() 创建切片函数
   	[]int 创建切片类型
    5 切片长度，此时没有显示指定该切片的容量，默认为5

   slice1 :=make([]int 4,10)

	容量必须>=长度，我们是不能创建长度大于容量的切片的。

	强调一下，。数组和切片的区别
	数组	array :=[2]int{12,43}
	切片	slice :=[]int{3:4}
*/

func main_01() {
	slice := []int{3, 3, 45, 656, 5}
	newSlice := slice[1:3]
	newSlice[0] = 133
	fmt.Println(slice)
	fmt.Printf("slice的类型是:%T", slice)
	fmt.Println(newSlice)
	//fmt.Printf("newSlice的类型是：%T",newSlice)
	/*
		给予原数组或者切片创建一个新的切片后，新的切片的容量和长度

		对于底层数组容量是k的切片slice[i:j]来说
		长度：j-i
		容量:k-i

	*/
	fmt.Printf("newSlice的长度是: %d\t,容量是：%d\n", len(newSlice), cap(newSlice)) /*
		输出
		newSlice的长度是: 2	,容量是：4
	*/
	//以上基于一个数组或者切片使用2个索引创建新切片的方法，此外还有一种3个索引的方法，
}

/*第3个用来限定新切片的容量，其用法为slice[i:j:k]。*/
func main_02() {

	slice1 := []int{1, 2, 3, 4, 5}
	newSlice := slice1[1:2:3]
	fmt.Printf("newSlice的长度为: %d\t,容量为：%d\n", len(newSlice), cap(newSlice)) /*
		输出
		newSlice的长度为: 1	,容量为：2
	*/
	//以上基于一个数组或者切片使用2个索引创建新切片的方法，此外还有一种3个索引的方法，第3个用来限定新切片的容量，其用法为slice[i:j:k]。
}

/*使用切片*/
func main_04() {

	slice := []string{"北京", "上海", "兰州", "南京"}
	for key, vlus := range slice {
		fmt.Printf("%d = %s\n", key, vlus)
	}
	//获取数组长度
	fmt.Println("slice数组长度为：", len(slice))
	//获取值
	fmt.Println(slice[0])
	//修改值
	slice[3] = "甘肃"
	fmt.Println(slice[3])

	slice_01 := []string{"北京", "上海", "兰州", "南京", "吉林", "长春"}
	newSlice_01 := slice_01[:3] //out [北京 上海 兰州]
	newSlice_02 := slice_01[:]  //out  [北京 上海 兰州 南京 吉林 长春]
	fmt.Println(newSlice_01)
	fmt.Println(newSlice_02)
}

func main_05() {
	//使用append 函数为新创建切片newSlice_01添加元素
	slice_01 := []string{"北京", "上海", "兰州", "南京", "吉林", "长春"} //数组
	newSlice_01 := slice_01[:]                               //创建一个新的切片，并且，切片长度为len+1--->len-1
	//使用append函数向新创建的newSlice_01中添加一个元素
	newSlice_01 = append(slice_01, "辽宁")
	fmt.Println(newSlice_01) //out [北京 上海 兰州 辽宁]
	fmt.Printf("新切片newSlice_01的类型为: %T\n", newSlice_01)
	fmt.Printf("切片newSlice_01的长度为：%d\n", len(newSlice_01))
	/*
		内置的append也是一个可变参数的函数，所以我们可以同时追加好几个值。
	*/
	fmt.Println("添加前", len(newSlice_01))
	newSlice_01 = append(slice_01, "内蒙古", "西藏")
	fmt.Println(newSlice_01)
	//变量新加元素切片newSlice_01
	for key, ok := range newSlice_01 {
		fmt.Println(key, ok)
	}
	fmt.Println("添加后", len(newSlice_01))
}

/*通过...操作符，把一个切片追加到另一个切片里。*/
func main_06() {
	slcie_inf := []string{"北京", "上海", "兰州"}
	new_Slice_inf := slcie_inf[:2]
	fmt.Println("添加前", new_Slice_inf)
	new_Slice_inf = append(new_Slice_inf, slcie_inf...)
	fmt.Println("添加后", new_Slice_inf)
}

func main_001() {

	slice_02 := make([]string, 4, 7)
	fmt.Println(len(slice_02))
	//slice_01[0] = "北京"
	slice_02 = append(slice_02, "上海", "南京", "甘肃", "广州", "云南", "贵州", "江西", "南宁")
	for i, s := range slice_02 {
		fmt.Println(i, s)
	}
	fmt.Println(slice_02, len(slice_02))
}

func main() {
	slice := make([]string, 1)
	slice = append(slice, "a", "b", "c")
	for i, s := range slice {
		fmt.Println(i, s)
	}
}
