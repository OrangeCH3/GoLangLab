/*
@Time : 2021/7/22 19:59
@Author : OrangeCH3
@Software: GoLand
@File : main
*/

package main

import (
	"fmt"
	"math/rand"
)

// 条件判断
func main() {

	age := rand.Int()
	if age > 18 {
		fmt.Println("您已成年，可以进入!")
	} else {
		fmt.Println("未成年禁止入内!")
	}

	// 多个条件进行判断
	if age > 35 {
		fmt.Println("人到中年，身不由己!")
	} else if age > 18 {
		fmt.Println("身体是革命的本钱!")
	} else {
		fmt.Println("好好学习，天天向上!")
	}
}
