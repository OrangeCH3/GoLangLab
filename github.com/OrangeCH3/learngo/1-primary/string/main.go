/*
@Time : 2021/7/22 15:58
@Author : OrangeCH3
@Software: GoLand
@File : main
*/

package main

import (
	"fmt"
	"strings"
)

// 定义多行字符串
var strMultiLine = `
TO
GET
HER
Together
`

// Go中字符串必须用 "" 包裹起来，
// Go中字符用 '' 包裹
func main() {

	// 转义字符的使用
	fmt.Println("str := \"D:\\Code\\learngo\\go.exe\"")

	strTest := "I'm OK"
	fmt.Println(strTest)

	// 打印多行字符
	fmt.Print(strMultiLine)
	fmt.Println()

	// 字符串相关操作
	fmt.Println(len(strTest))

	strTestDitto := ", OrangeCH3!"
	concatenatedString := fmt.Sprintf("%s%s", strTest, strTestDitto)
	fmt.Println(concatenatedString)

	ret := strings.Split(concatenatedString, ", ")
	fmt.Println(ret)

	fmt.Println(strings.Contains(strTest, "OK"))

	fmt.Println(strings.HasPrefix(strTest, "I"))
	fmt.Println(strings.HasSuffix(strTest, "K"))

	strIndexTest := "OrangeCH3"
	fmt.Println(strings.Index(strIndexTest, "C"))
	fmt.Println(strings.LastIndex(strIndexTest, "H"))

	fmt.Println(strings.Join(ret, "+"))
}
