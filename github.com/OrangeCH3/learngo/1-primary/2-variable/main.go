/*
@Time : 2021/7/21 16:04
@Author : OrangeCH3
@Software: GoLand
@File : main
*/

package main

import "fmt"

// Go语言中推荐使用驼峰式的命名
// var studentName 6-string

// 声明变量
// var name 6-string
// var age int
// var isOk 7-bool

// 批量声明变量
var (
	name string // ""
	age  int    // 0
	isOk bool   // false
)

func main() {

	// Go语言中非全局变量声明必须使用，否则无法编译

	name = "OrangeCH3"
	age = 17
	isOk = true

	// 全局变量声明赋值后可不使用
	fmt.Printf("Name: %s | Age: %d | Sex: (Default: Boy) %t\n", name, age, isOk) // %s是占位符
	fmt.Print("End\n")
	fmt.Println()

	// 声明变量的同时进行赋值
	var ultraman string = "mebius"
	fmt.Printf("Ultraman: %s\n", ultraman)

	// 简短变量声明，只能在函数中使用
	ultra := "seven"
	fmt.Printf("Ultra: %s\n", ultra)
}
