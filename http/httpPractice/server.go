package main

import (
	"fmt"
	"io"
	"net/http"
)

func Server() {
	serverProcess()

	http.ListenAndServe("127.0.0.1:8080", nil)
}

func serverProcess() {
	http.HandleFunc("/hello", SayHello)
	http.HandleFunc("/birthday", happyBirthday)
}

// 输入指定的网址, 显示hello http,hello world
func SayHello(w http.ResponseWriter, r *http.Request) {
	n, err := io.WriteString(w, "hello http,hello world")
	if err != nil {
		fmt.Println("sayhello failed")
	}

	fmt.Println("n = ", n)
}

// 输入指定的网址, 显示hello http,hello world
func happyBirthday(w http.ResponseWriter, r *http.Request) {
	n, err := io.WriteString(w, "honey, happy birthday!")
	if err != nil {
		fmt.Println("birthday failed")
	}

	fmt.Println("n = ", n)
}

func main() {
	Server()

}
