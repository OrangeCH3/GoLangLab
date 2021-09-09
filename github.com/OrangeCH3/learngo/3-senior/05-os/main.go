/*
@Time : 2021/9/8 9:54
@Author : OrangeCH3
@Software: GoLand
@File : main
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

// 基本方式读取文件
func basicReadIO() {
	// 只读方式打开当前目录下的ditto.txt文件
	file, err := os.Open("./github.com/OrangeCH3/learngo/3-senior/05-os/ditto.txt")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
	// 使用Read方法读取数据
	var tmp = make([]byte, 128)
	n, err := file.Read(tmp)
	if err == io.EOF {
		fmt.Println("文件读完了")
		return
	}
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Printf("读取了%d字节数据\n", n)
	fmt.Println(string(tmp[:n]))

}

// for循环读取文件
func forReadIO() {

	// 只读方式打开当前目录下的ditto.txt文件
	file1, err1 := os.Open("./github.com/OrangeCH3/learngo/3-senior/05-os/ditto.txt")
	if err1 != nil {
		fmt.Println("open file failed!, err:", err1)
		return
	}
	defer func(file1 *os.File) {
		err := file1.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file1)
	// 循环读取文件
	var content []byte
	var tmp1 = make([]byte, 128)
	for {
		n, err := file1.Read(tmp1)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		content = append(content, tmp1[:n]...)
	}
	fmt.Println(string(content))
}

// bufio按行读取示例
func bufReadIO() {
	file, err := os.Open("./github.com/OrangeCH3/learngo/3-senior/05-os/ditto.txt")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') //注意是字符
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Print(line)
	}
}

// ioutil.ReadFile读取整个文件
func ioutilReadIO() {
	content, err := ioutil.ReadFile("./github.com/OrangeCH3/learngo/3-senior/05-os/ditto.txt")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content))
}

// Write和WriteString
func basicWriteIO() {
	file, err := os.OpenFile("./github.com/OrangeCH3/learngo/3-senior/05-os/wrditto.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
	str := "hello, OrangeCH3!\n"
	file.Write([]byte(str))         //写入字节切片数据
	file.WriteString("Hello, Sun!") //直接写入字符串数据
}

// bufio.NewWriter
func bufferWriteIO() {
	file, err := os.OpenFile("./github.com/OrangeCH3/learngo/3-senior/05-os/wrbuf.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("Tick: ") //将数据先写入缓存
		writer.WriteString(strconv.Itoa(i))
		writer.WriteString("\n")
	}
	writer.Flush() //将缓存中的内容写入文件
}

// ioutil.WriteFile
func main() {
	str := "Hello OrangeCH3"
	err := ioutil.WriteFile("./github.com/OrangeCH3/learngo/3-senior/05-os/wrioutil.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}
