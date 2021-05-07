package main

import "fmt"

func printArray2(myArray []int) {
	// 引用传递
	// _表示匿名的变量
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}

	myArray[0] = 100
}

func main() {
	myArray := []int{1, 2, 3, 4} // 动态数组，切片 slice

	fmt.Printf("myArray type is %T\n", myArray)

	printArray2(myArray)

	fmt.Println("====")

	for _, value := range myArray {
		fmt.Println("value = ", value)
	}

}

/*
期望的结果：

myArray type is []int
value =  1
value =  2
value =  3
value =  4
====
value =  100
value =  2
value =  3
value =  4
*/

/*
# 数组 特征：

- 固定的：数组的长度是固定的；

- 数组类型：固定长度的数组在传参的时候，是严格匹配数组类型的

- 值拷贝
func printArray(myArray [4]int) {
	//值拷贝


# slice（动态数组）特征：

- 动态数组在传参上是引用传递，而且不同的元素长度的动态数组，他们的形参是一致

myArray := []int{1,2,3,4} // 动态数组，切片 slice

func printArray(myArray []int) {
	//引用传递
*/
