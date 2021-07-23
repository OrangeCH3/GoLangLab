/*
@Time : 2021/7/20 22:32
@Author : OrangeCH3
@Software: GoLand
@File : main
*/

// go build -o xxx.exe 指定名称构建可执行文件

package main

// 导入语句
import "fmt"

// 函数外部只能放置标识符（变量/常量/函数/类型）的声明
// fmt.Println("OrangeCH3, Let's go!") 非法

// 程序的入口函数
func main() {
	fmt.Println("OrangeCH3, Let's go!")
}
