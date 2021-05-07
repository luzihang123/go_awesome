package main

import "fmt"

/*
func swap(a int, b int) {
	// 只在内部交换，并不会在main中交换值
	var temp int
	temp = a
	a = b
	b = temp
}
*/

func swap(pa *int, pb *int) {
	var temp int
	temp = *pa // temp = main::a
	*pa = *pb  // main::a = main::b
	*pb = temp // main::b = temp
}
func main() {
	var a int = 10
	var b int = 20

	//swap(a, b)
	//a, b = b, a 这种可以交换

	swap(&a, &b)

	fmt.Println("a = ", a, " b = ", b)

	var p *int

	p = &a

	fmt.Println(&a)
	fmt.Println(p)

	var pp **int // 二级指针

	pp = &p

	fmt.Println(&p)
	fmt.Println(pp)
}
