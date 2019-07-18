package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	GetTest()
}

func GetTest() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}

	// 当获取到响应体的时候 一定要手动关闭 连接.
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("ioutil.ReadAll failed")
	}

	fmt.Println("body = ", string(body))
}
