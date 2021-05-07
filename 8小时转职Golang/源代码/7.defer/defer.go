package main

import "fmt"

func main() {
	// 写入defer关键字
	// Defer 用来保证一个函数调用会在程序执行的最后被调用。通常用于资源清理工作。
	// 遇见defer依次入栈，函数返回依次出栈执行
	defer fmt.Println("main end1")
	defer fmt.Println("main end2")

	fmt.Println("main::hello go 1")
	fmt.Println("main::hello go 2")
}
