/*
@Time : 2021/8/14 15:27
@Author : OrangeCH3
@Software: GoLand
@File : main
*/

package main

import (
	"errors"
	"fmt"
	"strings"
)

// 函数的定义
func intSum(x int, y int) int {
	return x + y
}

// 可变参数函数的定义
func intSum2(x ...int) int {
	// fmt.Println(x) //x是一个切片
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

// 多返回值函数的定义
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

//定义全局变量num
var num int64 = 10

// 包含全局变量函数的定义
func testGlobalVar() {
	fmt.Printf("num=%d\n", num) //函数中可以访问全局变量num
}

// 函数类型的定义
type calculation func(int, int) int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

// 函数作为参数的定义
func addD(x, y int) int {
	return x + y
}
func calcD(x, y int, op func(int, int) int) int {
	return op(x, y)
}

// 函数作为返回值
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

// 闭包函数的定义
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// defer函数的经典案例
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

// panic/recover函数的案例
func funcA() {
	fmt.Println("func A")
}

func funcB() {
	defer func() {
		err := recover()
		//如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}

func main() {

	// 函数的调用
	ret := intSum(10, 20)
	fmt.Println(ret) // 30

	// 可变参数函数的调用
	ret1 := intSum2()
	ret2 := intSum2(10)
	ret3 := intSum2(10, 20)
	ret4 := intSum2(10, 20, 30)
	fmt.Println(ret1, ret2, ret3, ret4) // 0 10 30 60

	// 多返回值函数的调用
	ret5, ret6 := calc(10, 2)
	fmt.Println(ret5, ret6)

	// 包含全局变量函数的调用
	testGlobalVar() //num=10

	// 函数类型变量
	var c calculation               // 声明一个calculation类型的变量c
	c = add                         // 把add赋值给c
	fmt.Printf("type of c:%T\n", c) // type of c:main.calculation
	fmt.Println(c(1, 2))            // 像调用add一样调用c

	f := add                        // 将函数add赋值给变量f1
	fmt.Printf("type of f:%T\n", f) // type of f:func(int, int) int
	fmt.Println(f(10, 20))          // 像调用add一样调用f

	// 函数作为参数的调用
	ret7 := calcD(10, 20, addD)
	fmt.Println(ret7) //30

	// 函数作为返回值的调用
	funcS, _ := do("+")
	ret8 := funcS(1, 6)
	print(ret8)

	// 闭包函数的调用
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg
	fmt.Println(txtFunc("test")) //test.txt

	// defer函数的调用
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())

	// panic/recover的运行环境
	funcA()
	funcB()
	funcC()
}
