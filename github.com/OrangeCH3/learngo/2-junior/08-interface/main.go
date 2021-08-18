/*
@Time : 2021/8/17 16:36
@Author : OrangeCH3
@Software: GoLand
@File : main
*/

package main

import "fmt"

// Sayer 接口
type Sayer interface {
	say()
}

type Mover interface {
	move()
}

type dog struct{}

type cat struct{}

// dog实现了Sayer接口
func (d dog) say() {
	fmt.Println("汪汪汪")
}

// dog实现了Mover接口
func (d dog) move() {
	fmt.Println("狗会动")
}

// cat实现了Sayer接口
func (c cat) say() {
	fmt.Println("喵喵喵")
}

// <-------我是一条分割线------->

// 类型断言判断函数
func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}

func main() {

	var x Sayer // 声明一个Sayer类型的变量x
	a := cat{}  // 实例化一个cat
	b := dog{}  // 实例化一个dog
	x = a       // 可以把cat实例直接赋值给x
	x.say()     // 喵喵喵
	x = b       // 可以把dog实例直接赋值给x
	x.say()     // 汪汪汪

	// 使用值接收者实现接口之后，不管是dog结构体还是结构体指针*dog类型的变量都可以赋值给该接口变量
	var y Mover
	var wangcai = dog{} // 旺财是dog类型
	y = wangcai         // x可以接收dog类型
	y.move()
	var fugui = &dog{} // 富贵是*dog类型
	y = fugui          // x可以接收*dog类型
	y.move()

	// 定义一个空接口x
	var z interface{}
	s := "Hello OrangeCH3"
	z = s
	fmt.Printf("type:%T value:%v\n", z, z)
	i := 100
	z = i
	fmt.Printf("type:%T value:%v\n", z, z)
	c := true
	z = c
	fmt.Printf("type:%T value:%v\n", z, z)

	// 空接口作为map值
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "OrangeCH3"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)

	// <-------我是一条分割线------->

	// 类型断言函数的调用
	var d interface{}
	d = "Hello OrangeCH3"
	justifyType(d)
}
