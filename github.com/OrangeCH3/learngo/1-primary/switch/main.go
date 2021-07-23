/*
@Time : 2021/7/23 17:22
@Author : OrangeCH3
@Software: GoLand
@File : main
*/

package main

import "fmt"

// 用 switch 简化大量的判断
func main() {
	finger := 4
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入！")
	}
	fmt.Println()

	// switch变种
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}
}
