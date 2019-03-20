package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"

	message "github.com/zxccl0518/go_study/chatroom/common/message"
)

func readPkg(conn net.Conn, mes *message.Message) error {
	buf := make([]byte, 1034*8)

	n, err := conn.Read(buf)
	if n != 4 || err != nil {
		fmt.Println("conn.Read() failed, err =", err)
		return err
	}

	var pkgLen = binary.BigEndian.Uint32(buf[:n])
	fmt.Println("客户端发送过来的消息 长度为 = ", int(pkgLen))
	fmt.Println("客户端发送过来的关于长度的 消息 内容为 = ", buf[:n])

	n, err = conn.Read(buf)
	if err != nil || n != int(pkgLen) {
		fmt.Println("服务器接收端，接收到的消息不完整。 err = ", err)
		return err
	}

	fmt.Println("服务器接收到 登录消息 = ", string(buf[:n]))
	err = json.Unmarshal(buf[:n], mes)
	if err != nil {
		fmt.Println("反序列化 失败 err = ", err)
		return err
	}

	return nil
}

func process(conn net.Conn) {
	defer conn.Close()

	for {
		var mes message.Message
		err := readPkg(conn, &mes)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端断开连接---")
				return
			} else {
				fmt.Println("读取 客户端 传过来的登陆消息 失败 err = ", err)
			}
		} else {
			fmt.Println("读取 客户端发送过来的 登录消息 成功。")
		}
	}
}

func main() {
	// 启动服务器
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("listen() start faild ,err = ", err)
		return
	}
	defer listen.Close()

	for {
		fmt.Println("等待 客户端的连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("服务器 监听客户端的连接失败 err = ", err)
		}

		go process(conn)
	}
}
