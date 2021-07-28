/*
@Time : 2021/7/23 18:07
@Author : OrangeCH3
@Software: GoLand
@File : main
*/

package main

import "fmt"

// 数组
// 必须指定存放元素的类型和个数
func main() {

	var testArray [3]int                        // 数组会初始化为int类型的零值
	var numArray = [3]int{1, 2, 7}              // 使用指定的初始值完成初始化
	var cityArray = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化
	fmt.Println(testArray)                      // [0 0 0]
	fmt.Println(numArray)                       // [1 2 0]
	fmt.Println(cityArray)                      // [北京 上海 深圳]

	// 自动推断赋值[...]
	testDitto := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(testDitto)

	// 根据索引来初始化
	testDittoo := [5]int{0: 1, 4: 7}
	fmt.Println(testDittoo)

	// 数组的遍历
	for i := 0; i < len(cityArray); i++ {
		fmt.Println(cityArray[i])
	}
	for i, v := range cityArray {
		fmt.Println(i, v)
	}

	// 多维数组
	multidimensionalArrays := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	for _, v1 := range multidimensionalArrays {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}
}
