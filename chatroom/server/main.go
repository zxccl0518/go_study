package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"

	"github.com/zxccl0518/go_study/chatroom/common/message"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 1034*8)

	// conn.Read在conn没有关闭的情况下，才会阻塞。
	// 如果客户端关闭了，conn则不会阻塞
	n, err := conn.Read(buf)
	if n != 4 || err != nil {
		fmt.Println("conn.Read() failed, err =", err)
		return
	}

	var pkgLen = binary.BigEndian.Uint32(buf[:n])
	fmt.Println("客户端发送过来的消息 长度为 = ", int(pkgLen))
	fmt.Println("客户端发送过来的关于长度的 消息 内容为 = ", buf[:n])

	n, err = conn.Read(buf)
	if err != nil || n != int(pkgLen) {
		fmt.Println("服务器接收端，接收到的消息不完整。 err = ", err)
		return
	}

	// fmt.Println("服务器接收到 登录消息 打印的是字节切片形式的内容 = ", buf[:n])
	fmt.Println("服务器接收到 登录消息 打印的是string形式的内容 = ", string(buf[:n]))
	// 计数就是一层窗户纸。 不弄清楚的话 永远都是个谜题。弄清楚了就如同 捅破窗户纸一般 简单。
	err = json.Unmarshal(buf[:n], &mes) // 这个位置一定要用 &mes 才会将序列化的内容反序列化到mes中。
	if err != nil {
		fmt.Println("反序列化 失败 err = ", err)
		return
	}

	return
}

func writePkg(conn net.Conn) {
	var loginResMes message.LoginResMessage
	loginResMes.Code = 200
	loginResMes.Err = nil

	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("序列化 失败 err = ", err)
		return
	}

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("服务器 返回给客户端 消息失败。。。")
		return
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		// 这里我们阿静读取数据包，直接封装成一个函数readPkg，返回Message,Err
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端断开连接---")
				return
			}

			fmt.Println("读取客户端 传过来的登陆消息失败，c错误原因非客户端掉线。 err = ", err)
			return
		}

		fmt.Println("读取 客户端发送过来的 登录消息 成功。")
		mes.Print()

		// 接下来，服务器给客户端发送一个消息，通知登录情况。
		writePkg(conn)
		return
	}
}

func main() {
	// 启动服务器
	listen, err := net.Listen("tcp", "localhost:8887")
	if err != nil {
		fmt.Println("listen() start faild ,err = ", err)
		return
	}
	defer listen.Close()
	fmt.Println("服务器端启动成功...")

	for {
		fmt.Println("循环阻塞，等待 客户端的连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("服务器 监听客户端的连接失败 err = ", err)
		}

		go process(conn)
	}
}
