package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// 定义一个全局变量。
var (
	pool *redis.Pool
)

// 启动程序的时候，就会初始化连接池。
func init() {
	pool = &redis.Pool{
		MaxIdle:     9,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	conn := pool.Get()
	defer conn.Close()

	// 设置一个 值
	_, err := conn.Do("set", "name", "tom猫")
	if err != nil {
		fmt.Println("conn Do err = ", err)
		return
	}

	// 去除一个值
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("conn Get err = ", err)
		return
	}

	fmt.Println("r = ", r)
}
