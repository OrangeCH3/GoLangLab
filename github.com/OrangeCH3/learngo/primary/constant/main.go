/*
@Time : 2021/7/21 16:51
@Author : OrangeCH3
@Software: GoLand
@File : main
*/

package main

import "fmt"

// 常量
const pi = 3.1415926

// 批量声明常量
const (
	statusOk = 200
	notFound = 404
)

// 批量声明常量时，若未声明则默认和上一行同值
// 特殊地：iota 会随行递增，可以理解为 const 语句块的行索引
const (
	tap        = iota // 0
	tapDitto          // 1
	tapDittoo         // 2
	tapDittooo        // 3
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {

	fmt.Println("PI:", pi)
	fmt.Println("StatusOk:", statusOk)
	fmt.Println("NotFound:", notFound)

	fmt.Println()
	fmt.Println(tap, tapDitto, tapDittoo, tapDittooo)

	fmt.Println()
	fmt.Println(KB, MB, GB, TB, PB)

}
