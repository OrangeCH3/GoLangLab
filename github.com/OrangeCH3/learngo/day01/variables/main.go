/*
@Time : 2021/7/21 16:04
@Author : OrangeCH3
@Software: GoLand
@File : main
*/

package main

import "fmt"

// Go语言中推荐使用驼峰式的命名
// var studentName string

// 声明变量
// var name string
// var age int
// var isOk bool

// 批量声明变量
var (
	name string // ""
	age  int    // 0
	isOk bool   // false
)

func main() {

	name = "OrangeCH3"
	age = 17
	isOk = true

	// Go语言中变量声明必须使用，否则无法编译
	fmt.Printf("Name: %s | Age: %d | Sex: (Default: Boy) %t\n", name, age, isOk) // %s是占位符
	fmt.Println()
	fmt.Print("End")
}
