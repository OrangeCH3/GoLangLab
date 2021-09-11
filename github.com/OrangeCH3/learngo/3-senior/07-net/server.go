/*
@Time : 2021/9/10 11:10
@Author : OrangeCH3
@Software: GoLand
@File : server
*/

package main

import (
	"fmt"
	"net/http"
)

// http server

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, OrangeCH3!")
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}
