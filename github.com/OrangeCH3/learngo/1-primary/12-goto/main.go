/*
@Time : 2021/7/23 17:30
@Author : OrangeCH3
@Software: GoLand
@File : main
*/

package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	// label标签
breakTag:
	fmt.Println("结束for循环")
}
