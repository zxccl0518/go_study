package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	// 循环的接受 客户端发过来的数据。
	defer conn.Close() // 关闭 conn 不关闭的话，会影响其他客户端连接服务器的性能
	for {
		// 创建一个 新的切片
		buf := make([]byte, 1024)
		// conn.Read(buf)
		// 等待客户端通过conn发送消息，
		// 如果客户端没有write[发送] ，那么协程就会阻塞在这里
		clientAddr := conn.RemoteAddr().String()
		// fmt.Printf("服务器 在等待 %v 客户端 发送的消息。", clientAddr)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("客户端已经退出。 err = ", err)
			return
		}
		str := string(buf[:n])
		if str == "exit" {
			fmt.Println("服务器端接收到信息，客户端请求退出。")
			break
		} else {
			// 显示客户端发送的内容的到服务器的终端
			fmt.Println("clientAddr[" + clientAddr + "] : " + str)
		}

	}
}

func main() {
	fmt.Println("服务器开始监听 ...")
	// net.Listen("tcp", "127.0.0.1:8888")	// 这种写法 支持iPv4 不支持ipv6
	listen, err := net.Listen("tcp", "0.0.0.0:8889") // 这种写法 支持iPv4 也支持ipv6
	if err != nil {
		fmt.Println("监听失败 err = ", err)
		return
	}
	defer listen.Close() // 延时关闭。

	// 循环等待 客户端 来链接我。
	for {
		// 等待客户端连接。
		fmt.Println("等待 客户端的额连接。。。")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err = ", err)
			return
		} else {
			fmt.Printf("Accept() suc con = %v 客户端ip地址：%v\n", conn, conn.RemoteAddr().String())
		}

		// 这里准备起一个协程，为客户端服务。。
		process(conn)
	}
	// fmt.Printf("listen suc = %v \n ", listen)
}
