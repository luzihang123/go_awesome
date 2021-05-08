// 切片的截取
package main

import "fmt"

func main() {
	/* 创建切片 */
	s := []int{1, 2, 3} // len = 3, cap = 3, [1,2,3]

	/* 打印原始切⽚ */
	fmt.Println("打印原始切⽚ s ==", s)

	// [0, 2] 左闭右开 打印⼦子切片从索引0(包含) 到索引2(不包含)
	s1 := s[0:2] // [1, 2]

	fmt.Println("s1 = ", s1)

	s1[0] = 100

	fmt.Println("s = ", s)
	fmt.Println("s1 = ", s1)

	// copy可以将底层数组的slice一起进行拷贝
	s2 := make([]int, 3) // s2 = [0,0,0]

	// 将s中的值 依次拷贝到s2中
	copy(s2, s)
	fmt.Println("s2 = ", s2)

}
