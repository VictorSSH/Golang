package main

import (
	f "fmt"
	"math/rand"
	"time"
)

func main() {
	/*求出一个数组的最大值，并得到该数组最大值的下标*/
	var intArray [5]int = [...]int{1, -1, 9, 99, 34}
	maxVal := intArray[0]
	maxValIndex := 0
	for i := 0; i < len(intArray); i++ {
		//从第二个元素开始循环比较。如果发现有更大的，则交换
		if maxVal < intArray[i] {
			maxVal = intArray[i]
			maxValIndex = i
		}
	}
	f.Printf("maxVal=%v maxValIndex=%v\n", maxVal, maxValIndex)

	/*求出一个数组的平均值，for-range*/

	//1，声明一个数组

	var arr01 [5]int = [...]int{4, 3, 2, 2, 5}
	sum := 0
	//2，求出和sum
	for _, val := range arr01 {
		//累加求和
		sum += val
	}

	//2, 求出平均值
	f.Printf("sum = %v 平均值= %v\n\n", sum, float64(sum)/float64(len(arr01)))

	/*要求：随机生成5个数，并将其反转打印*/
	var intArr02 [5]int
	//为了每次生成的随机数不一样，我们需要给一个seed值
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(intArr02); i++ {
		intArr02[i] = rand.Intn(100)
	}
	f.Println(intArr02)

}
