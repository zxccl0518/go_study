package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.1.100:8889")
	if err != nil {
		fmt.Println("client dial err = ", err)
		return
	}

	fmt.Println("conn 成功 = ", conn)

	// 功能1： 客户端可以发送单行数据，然后就退出。
	reader := bufio.NewReader(os.Stdin)

	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("从 肩旁读取内容失败 err = ", err)
			return
		}
		str = strings.Trim(str, " \r\n")

		if str == "exit" {
			fmt.Println("客戶端退出")
			break
		}

		n, err := conn.Write([]byte(str))
		if err != nil {
			fmt.Println("conn.write err = ", err)
		}
		fmt.Printf("客户端发送了 %d 字节的数据，并退出\n", n)
	}
}
