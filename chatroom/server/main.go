package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"

	"github.com/zxccl0518/go_study/chatroom/common/message"
)

func process(conn net.Conn) {
	// 读取客户端发送的信息。
	// 延迟关闭conn
	// defer conn.Close()

	// // 循环读取客户端的消息。
	// var recieveBuf = make([]byte, 1024*8)
	// // 读操作，如果读取不到，会阻塞在这里。
	// n, err := conn.Read(recieveBuf)
	// if err == io.EOF {
	// 	fmt.Println("客户端发送过来的数据 已经读取完毕。")
	// 	return
	// }
	// if err != nil || n != 4 {
	// 	fmt.Println("服务器 接受数据出错, err = ", err)
	// 	return
	// }
	// fmt.Println("服务器读取到的内容长度 recieveBuf len = ", string(recieveBuf[:n]))
	// var mesLen = int(binary.BigEndian.Uint32(recieveBuf[:n]))
	// fmt.Println("客户端第一次发送过来的数据是 长度， len = ", mesLen)

	// n, err = conn.Read(recieveBuf)
	// if err != nil || n != int(mesLen) {
	// 	fmt.Println("服务器接受 登录的数据长度不正确。err = ", err)
	// 	return
	// }

	// fmt.Println("服务器读取到的内容 recieveBuf = ", string(recieveBuf[:n]))
	// var mesStruct message.Message
	// err = json.Unmarshal(recieveBuf[:n], &mesStruct)
	// if err != nil {
	// 	fmt.Println("服务器端 反序列化失败，err = ", err)
	// 	return
	// }

	// fmt.Println("返 序列化 成功")
	// mesStruct.PrintMessage()

	buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据...")
	//conn.Read 在conn没有被关闭的情况下，才会阻塞
	//如果客户端关闭了 conn 则，就不会阻塞
	_, err := conn.Read(buf[:4])
	if err != nil {
		err = errors.New("read pkg header error")
		return
	}
	//根据buf[:4] 转成一个 uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])
	//根据 pkgLen 读取消息内容
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		err = errors.New("read pkg body error")
		return
	}
	//把pkgLen 反序列化成 -> message.Message
	// 技术就是一层窗户纸 &mes！！
	var mes message.Message
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarsha err=", err)
		return
	}
	return
}

func main() {
	// 提示信息。
	fmt.Println("服务器 在8889端口监听。")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net. listen err = ", err)
		return
	}

	// 一旦监听成功，就循环等待客户端来连接服务器。
	for {
		fmt.Println("等待客户端来连接服务器。。。。")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept() err = ", err)
		}

		// 一旦连接成功，就启动一个协程 和客户端保持通讯。
		go process(conn)
	}
}
