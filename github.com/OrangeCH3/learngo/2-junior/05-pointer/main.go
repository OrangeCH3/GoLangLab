/*
@Time : 2021/8/15 16:11
@Author : OrangeCH3
@Software: GoLand
@File : main
*/

package main

import "fmt"

// 指针传值函数定义
func modify1(x int) {
	x = 100
}
func modify2(x *int) {
	*x = 100
}

func main() {

	// 指针地址和指针类型
	a := 10
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc00001a078
	fmt.Printf("b:%p type:%T\n", b, b) // b:0xc00001a078 type:*int
	fmt.Println(&b)                    // 0xc00000e018

	// 指针取值
	fmt.Printf("type of b:%T\n", b)
	c := *b // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)

	// 指针传值函数调用
	modify1(a)
	fmt.Println(a) // 10
	modify2(&a)
	fmt.Println(a) // 100

	// new函数的使用
	var d *int
	// 指针作为引用类型需要初始化
	d = new(int)
	*d = 10
	fmt.Println(*d)

	// make函数的使用
	var e map[string]int
	e = make(map[string]int, 10)
	e["OrangeCH3"] = 100
	fmt.Println(e)
}
