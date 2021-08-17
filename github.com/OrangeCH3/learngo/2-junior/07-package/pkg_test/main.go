/*
@Time : 2021/8/17 10:29
@Author : OrangeCH3
@Software: GoLand
@File : main
*/

package pkg2

import "fmt"

// 导入的包会最先初始化并调用其init()函数
func init() {
	fmt.Println("This is pkg_test init() function!")
}

// 包变量可见性

var a = 100 // 首字母小写，外部包不可见，只能在当前包内使用

// Mode 首字母大写外部包可见，可在其他包中使用
const Mode = 36

type person struct { // 首字母小写，外部包不可见，只能在当前包内使用
	name string
}

// Add 首字母大写，外部包可见，可在其他包中使用
func Add(x, y int) int {
	return x + y
}

func age() { // 首字母小写，外部包不可见，只能在当前包内使用
	var Age = 18 // 函数局部变量，外部包不可见，只能在当前函数内使用
	fmt.Println(Age)
}
