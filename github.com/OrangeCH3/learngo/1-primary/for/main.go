/*
@Time : 2021/7/22 20:10
@Author : OrangeCH3
@Software: GoLand
@File : main
*/

package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println()

	// for循环变种1
	j := 7
	for ; j < 10; j++ {
		fmt.Println(j)
	}
	fmt.Println()

	// for循环变种2
	k := 6
	for k < 10 {
		fmt.Println(k)
		k++
	}

	// for range循环
	str := "OrangeCH3"
	for i, v := range str {
		fmt.Printf("Index:%d | Value:%c\n", i, v)
	}

}
