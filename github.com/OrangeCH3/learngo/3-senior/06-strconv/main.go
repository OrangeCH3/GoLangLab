/*
@Time : 2021/9/9 14:59
@Author : OrangeCH3
@Software: GoLand
@File : main
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Atoi()
	s1 := "100"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("can't convert to int")
	} else {
		fmt.Printf("type:%T value:%#v\n", i1, i1) //type:int value:100
	}

	// Itoa()
	i2 := 200
	s2 := strconv.Itoa(i2)
	fmt.Printf("type:%T value:%#v\n", s2, s2) //type:string value:"200"

	// ParseXxx()
	b, err := strconv.ParseBool("true")
	f, err := strconv.ParseFloat("3.1415", 64)
	i, err := strconv.ParseInt("-2", 10, 64)
	u, err := strconv.ParseUint("2", 10, 64)
	fmt.Println(b, f, i, u)

	// FormatXxx()
	s11 := strconv.FormatBool(true)
	s12 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	s13 := strconv.FormatInt(-2, 16)
	s14 := strconv.FormatUint(2, 16)
	fmt.Println(s11, s12, s13, s14)
}
