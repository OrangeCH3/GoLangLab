/*
@Time : 2021/7/21 20:17
@Author : OrangeCH3
@Software: GoLand
@File : main
*/

package main

import "fmt"

// 整型
func main() {

	// 类型推断，十进制
	var numOne = 10
	fmt.Printf("%d\n", numOne)
	fmt.Printf("%b\n", numOne) // 占位符%b表示二进制

	// 八进制
	numTwo := 077
	fmt.Printf("%d\n", numTwo)
	fmt.Printf("%o\n", numTwo) // 占位符%o表示八进制

	// 十六进制
	numThree := 0xffff
	fmt.Printf("%d\n", numThree)
	fmt.Printf("%x\n", numThree) // 占位符%x表示十六进制

	// 查看变量的类型
	fmt.Printf("%T\n", numThree)

}
