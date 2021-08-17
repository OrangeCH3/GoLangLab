/*
@Time : 2021/8/16 11:08
@Author : OrangeCH3
@Software: GoLand
@File : main
*/

package main

import (
	"encoding/json"
	"fmt"
)

// NewInt 类型定义
type NewInt int

// MyInt 类型别名
type MyInt = int

// Person 结构体
// 结构体中字段大写开头(Name)表示可公开访问，小写(name)表示私有（仅在定义当前结构体的包中可访问）
type Person struct {
	name string
	city string
	age  int8
}

// 构造函数
func newPerson(name, city string, age int8) *Person {
	return &Person{
		name: name,
		city: city,
		age:  age,
	}
}

//Dream Person做梦的方法
func (p Person) dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

// SetAge 设置p的年龄
// 使用指针接收者
func (p *Person) setAge(newAge int8) {
	p.age = newAge
}

// <-------我是一条分割线------->

// Address 地址结构体
// 结构体中字段大写开头(Province)表示可公开访问，小写(province)表示私有（仅在定义当前结构体的包中可访问）
type Address struct {
	Province string
	City     string
}

// User 用户结构体
// 结构体中字段大写开头(Name)表示可公开访问，小写(name)表示私有（仅在定义当前结构体的包中可访问）
type User struct {
	Name    string
	Gender  string
	Address Address
}

// <-------我是一条分割线------->

//Animal 动物
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

//Dog 狗
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

// <-------我是一条分割线------->

//Student 学生
type Student struct {
	ID     int
	Gender string
	Name   string
}

//Class 班级
type Class struct {
	Title    string
	Students []*Student
}

// <-------我是一条分割线------->

type PersonD struct {
	namee  string
	agee   int8
	dreams []string
}

// setDreams 使用传入的slice的拷贝进行结构体赋值
func (p *PersonD) setDreams(dreams []string) {
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)
}

func main() {

	// 1. NewInt-MyInt类型定义和类型别名的区别
	var a NewInt
	var b MyInt
	fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
	fmt.Printf("type of b:%T\n", b) //type of b:int

	// 2.1 结构体基本实例化及输出(%#v)
	var p1 Person
	p1.name = "OrangeCH3"
	p1.city = "北京"
	p1.age = 17
	fmt.Printf("p1=%v\n", p1)  //p1={OrangeCH3 北京 17}
	fmt.Printf("p1=%#v\n", p1) //p1=main.Person{name:"OrangeCH3", city:"北京", age:17}

	// 2.2 创建指针类型结构体实例化
	var p2 = new(Person)
	fmt.Printf("%T\n", p2) //*main.Person
	p2.name = "小王子"
	p2.age = 28
	p2.city = "上海"
	fmt.Printf("p2=%#v\n", p2) //p2=&main.Person{name:"小王子", city:"上海", age:28}

	// 2.3 取结构体的地址实例化
	p3 := &Person{}
	fmt.Printf("%T\n", p3)     //*main.Person
	fmt.Printf("p3=%#v\n", p3) //p3=&main.Person{name:"", city:"", age:0}
	p3.name = "七米"
	p3.age = 30
	p3.city = "成都"
	fmt.Printf("p3=%#v\n", p3) //p3=&main.Person{name:"七米", city:"成都", age:30}

	// 使用键值对初始化
	p4 := Person{
		name: "小王子",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p4=%#v\n", p4) //p4=main.Person{name:"小王子", city:"北京", age:18}

	// 调用构造函数
	p5 := newPerson("张三", "沙河", 90)
	fmt.Printf("%#v\n", p5) //&main.Person{name:"张三", city:"沙河", age:90}

	// 调用方法
	// 注意：方法属于特定的类型
	p5.dream()

	// 调用setAge方法
	p5.setAge(30)
	fmt.Println(p5.age) // 30

	// <-------我是一条分割线------->

	// 嵌套结构体的初始化及输出
	user1 := User{
		Name:   "小王子",
		Gender: "男",
		Address: Address{
			Province: "山东",
			City:     "威海",
		},
	}
	//user1=main.User{Name:"小王子", Gender:"男", Address:main.Address{Province:"山东", City:"威海"}}
	fmt.Printf("user1=%#v\n", user1)

	// <-------我是一条分割线------->

	// 结构体的“继承”
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{ //注意嵌套的是结构体指针
			name: "乐乐",
		},
	}
	d1.wang() //乐乐会汪汪汪~
	d1.move() //乐乐会动！

	// <-------我是一条分割线------->

	// 结构体与JSON序列化
	c := &Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu := &Student{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			ID:     i,
		}
		c.Students = append(c.Students, stu)
	}
	//JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)
	//JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)

	// <-------我是一条分割线------->

	// slice和map这两种数据类型都包含了指向底层数据的指针，因此我们在需要复制它们时要特别注意
	p11 := PersonD{namee: "小王子", agee: 18}
	dataa := []string{"吃饭", "睡觉", "打豆豆"}
	p11.setDreams(dataa)

	// 你真的想要修改 p11.dreams 吗？
	dataa[1] = "不睡觉"
	fmt.Println(p11.dreams) // 不会修改p11的值

	dataa[1] = "不睡觉"
	p11.setDreams(dataa)
	fmt.Println(p11.dreams) // 会修改p11的值

}
