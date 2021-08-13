/*
@Time : 2021/8/13 15:52
@Author : OrangeCH3
@Software: GoLand
@File : main
*/

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {

	// map初始化后再赋值
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of a:%T\n", scoreMap)

	// map声明时同时填充元素
	userInfo := map[string]string{
		"username": "OrangeCH3",
		"password": "123456",
	}
	fmt.Println(userInfo)

	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}

	// map的遍历过程
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	// map只遍历key
	for k := range scoreMap {
		fmt.Println(k)
	}

	// map删除指定的元素
	scoreMap["娜扎"] = 60
	delete(scoreMap, "娜扎") //将小明:100从map中删除
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	// 按照指定顺序遍历map
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMapSorted = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMapSorted[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMapSorted {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMapSorted[key])
	}

	// 元素为map类型的切片
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 1)
	mapSlice[1] = make(map[string]string, 1)
	mapSlice[2] = make(map[string]string, 1)
	mapSlice[0]["name"] = "OrangeCH3"
	mapSlice[1]["password"] = "123456"
	mapSlice[2]["address"] = "Earth"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}

	// 值为切片类型的map
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)

}
