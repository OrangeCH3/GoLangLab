/*
@Time : 2021/8/2 15:24
@Author : OrangeCH3
@Software: GoLand
@File : main
*/

package main

import "fmt"

func main() {

	// 声明切片类型
	var a []string              //声明一个字符串切片
	var b = []int{17, 36}       //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Println(a)              //[]
	fmt.Println(b)              //[17, 36]
	fmt.Println(c)              //[false true]
	fmt.Println(a == nil)       //true
	fmt.Println(b == nil)       //false
	fmt.Println(c == nil)       //false

	// 由数组得到切片
	ad := [5]int{1, 2, 3, 4, 5}
	s := ad[1:3] // s := a[low:high]
	// 容量从切片的第一个元素开始算起
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	// 切片再切片
	s2 := s[3:4] // 索引的上限是cap(s)而不是len(s)
	fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))

	// make()的用法
	am := make([]int, 2, 10)
	fmt.Println(am)      //[0 0]
	fmt.Println(len(am)) //2
	fmt.Println(cap(am)) //10

	// 切片不能直接比较
	var ss1 []int         //len(s1)=0;cap(s1)=0;s1==nil
	ss2 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
	ss3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil
	fmt.Println(ss1 == nil)
	fmt.Println(ss2 == nil)
	fmt.Println(ss3 == nil)
	fmt.Println(len(ss1))
	fmt.Println(len(ss2))
	fmt.Println(len(ss3))

	// 切片的拷贝问题
	s11 := make([]int, 3) //[0 0 0]
	s12 := s11            //将s1直接赋值给s2，s1和s2共用一个底层数组
	s12[0] = 100
	fmt.Println(s11) //[100 0 0]
	fmt.Println(s12) //[100 0 0]

	// 切片的遍历
	sl := []int{1, 3, 5}
	for i := 0; i < len(sl); i++ {
		fmt.Println(i, sl[i])
	}
	for index, value := range sl {
		fmt.Println(index, value)
	}

	// 切片添加元素
	var sa []int
	sa = append(sa, 1)       // [1]
	sa = append(sa, 2, 3, 4) // [1 2 3 4]
	sa2 := []int{5, 6, 7}
	sa = append(sa, sa2...) // [1 2 3 4 5 6 7]
	fmt.Println(sa)
	fmt.Println(sa2)

	//append()添加元素和切片扩容
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}

	// copy()复制切片
	ac := []int{1, 2, 3, 4, 5}
	cc := make([]int, 5, 5)
	copy(cc, ac)    //使用copy()函数将切片a中的元素复制到切片c
	fmt.Println(ac) //[1 2 3 4 5]
	fmt.Println(cc) //[1 2 3 4 5]
	cc[0] = 1000
	fmt.Println(ac) //[1 2 3 4 5]
	fmt.Println(cc) //[1000 2 3 4 5]

	// 从切片中删除元素
	ade := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// 要删除索引为2的元素
	ade = append(ade[:2], ade[3:]...)
	fmt.Println(ade) //[30 31 33 34 35 36 37]
}
