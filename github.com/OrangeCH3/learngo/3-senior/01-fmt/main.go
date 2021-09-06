/*
@Time : 2021/9/6 11:53
@Author : OrangeCH3
@Software: GoLand
@File : main
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	// Print
	fmt.Print("在终端打印该信息")
	name := "OrangeCH3"
	fmt.Printf("我是：%s\n", name)
	fmt.Println("在终端打印单独一行显示")

	// Fprint 向标准输出写入内容
	fprintln, err := fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	if err != nil {
		fmt.Println(fprintln)
		return
	}
	fileObj, err1 := os.OpenFile("./github.com/OrangeCH3/learngo/3-senior/01-fmt/demo.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err1 != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	namee := "OrangeCH3"
	// 向打开的文件句柄中写入内容
	fprintf, err2 := fmt.Fprintf(fileObj, "往文件中写如信息：%s", namee)
	if err2 != nil {
		fmt.Println(fprintf)
		return
	}

	// Sprint
	s1 := fmt.Sprint("->")
	nameee := "OrangeCH3"
	age := 17
	s2 := fmt.Sprintf("Name:%s,Age:%d", nameee, age)
	s3 := fmt.Sprintln("Name:Sun,Age:17")
	fmt.Println(s2, s1, s3)

	// Errorf
	err3 := fmt.Errorf("这是一个错误")
	fmt.Println(err3)

	// 通用占位符
	fmt.Printf("%v\n", 100)
	fmt.Printf("%v\n", false)
	o := struct{ name string }{"Sun"}
	fmt.Printf("%v\n", o)
	fmt.Printf("%#v\n", o)
	fmt.Printf("%T\n", o)
	fmt.Printf("100%%\n")

	// Scan
	//var (
	//	named    string
	//	aged     int
	//	married bool
	//)
	//scan, err4 := fmt.Scan(&named, &aged, &married)
	//if err4 != nil {
	//	fmt.Println(scan)
	//	return
	//}
	//fmt.Printf("扫描结果 name:%s age:%d married:%t \n", named, aged, married)

	// Scanf
	//var (
	//	name7    string
	//	age7     int
	//	married7 bool
	//)
	//_, err5 := fmt.Scanf("1:%s 2:%d 3:%t", &name7, &age7, &married7)
	//if err5 != nil {
	//	return
	//}
	//fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name7, age7, married7)

	// Scanln
	var (
		named   string
		aged    int
		married bool
	)
	_, err6 := fmt.Scanln(&named, &aged, &married)
	if err6 != nil {
		return
	}
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", named, aged, married)

}
