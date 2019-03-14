package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	// 通过 go 想redis 写入数据 和读取数据。
	// 1. 连接到redis

	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial() err = ", err)
		return
	}
	defer conn.Close()

	// fmt.Println("conn success ,,, conn = ", conn)
	// 通过 go 向 redis写入数据 string[key - val]
	_, err = conn.Do("Set", "name", "南风喃")
	if err != nil {
		fmt.Println("set redis err = ", err)
		return
	}

	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("get er = ", err)
		return
	}

	// 因为返回的r是interface{}
	// 因为 name 对应的值是string，因此我们需要转换
	// redis.String() redis 自带的方法已经提供 将interface{}类型的变量转换为 string类型

	fmt.Println("操作成功 r = ", r)
}
