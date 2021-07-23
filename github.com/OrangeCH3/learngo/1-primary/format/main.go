/*
@Time : 2021/7/22 15:50
@Author : OrangeCH3
@Software: GoLand
@File : main
*/

package main

import "fmt"

// fmt占位符
func main() {

	var num = 100
	// 查看类型
	fmt.Printf("%T\n", num)
	// 查看变量的值
	fmt.Printf("%v\n", num)
	// 查看各进制数值
	fmt.Printf("%b\n", num)
	fmt.Printf("%o\n", num)
	fmt.Printf("%d\n", num)
	fmt.Printf("%x\n", num)
	fmt.Println()

	var str = "Hello, OrangeCH3"
	fmt.Printf("%v\n", str)
	fmt.Printf("%s\n", str)
}
