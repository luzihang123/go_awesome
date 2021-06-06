// 反射
// 变量的内置pair结构详细说明
package main

import "fmt"

func main() {

	var a string
	// pair<statictype:,String value:"aceld">
	a = "aceld"

	// pair<type:string,value:"aceld">
	var allType interface{}
	allType = a

	str, _ := allType.(string)
	fmt.Println(str)
}
