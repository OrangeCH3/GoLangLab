/*
@Time : 2021/7/23 17:39
@Author : OrangeCH3
@Software: GoLand
@File : main
*/

package main

import "fmt"

func main() {
	a := 7
	b := 36
	if a < 9 && b > 29 {
		a <<= 1
		b >>= 1
		fmt.Println(a, b)
	}
}
