/*
@Time : 2021/9/7 18:47
@Author : OrangeCH3
@Software: GoLand
@File : main
*/

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.Println("这是一条很普通的日志。")
	v := "很普通的"
	log.Printf("这是一条%s日志。\n", v)
	//log.Fatalln("这是一条会触发Fatal的日志。")
	//log.Panicln("这是一条会触发Panic的日志。")

	// 配置logger
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")

	// 配置日志前缀
	log.SetFlags(log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("[OrangeCH3]")
	log.Println("这是一条很普通的日志。")

	// 配置日志输出位置
	logFile, err := os.OpenFile("./github.com/OrangeCH3/learngo/3-senior/04-log/log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("[OrangeCH3]")
	log.Println("这是一条很普通的日志。")

	// 创建自定义logger
	logger := log.New(os.Stdout, "[New-OrangeCH3]", log.Lshortfile|log.Ldate|log.Ltime)
	logger.Println("这是自定义的logger记录的日志。")
}
