package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

	"github.com/zxccl0518/go_study/chatroom/common/message"
)

func writePkg(conn net.Conn, data []byte) (err error) {

	var pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	// 发送长度
	_, err = conn.Write(buf[:4])
	if err != nil {
		fmt.Println("发送消息的长度失败 err  = ", err)
		return
	}

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("发送 消息本身失败 err = ", err)
		return
	}

	return
}

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 1024*8)

	// conn.Read在conn没有关闭的情况下，才会阻塞。
	// 如果客户端关闭了，conn则不会阻塞
	n, err := conn.Read(buf)
	if n != 4 || err != nil {
		fmt.Println("conn.Read() failed, err =", err)
		return
	}

	// 根据buf[:4] 转成一个uint32类型
	var pkgLen = binary.BigEndian.Uint32(buf[:n])
	fmt.Println("utils readPkg() 接收的消息 长度为 = ", int(pkgLen))
	fmt.Println("utils readPkg() 接收的消息 消息内容为 = ", buf[:n])

	// 根据pkgLen 读取消息内容
	n, err = conn.Read(buf[:int(pkgLen)])
	if err != nil || n != int(pkgLen) {
		fmt.Println("服务器接收端，接收到的消息不完整。 err = ", err)
		return
	}

	// fmt.Println("服务器接收到 登录消息 打印的是字节切片形式的内容 = ", buf[:n])
	fmt.Println("服务器接收到 登录消息 打印的是string形式的内容 = ", string(buf[:n]))
	// 技术就是一层窗户纸。 不弄清楚的话 永远都是个谜题。弄清楚了就如同 捅破窗户纸一般 简单。
	err = json.Unmarshal(buf[:n], &mes) // 这个位置一定要用 &mes 才会将序列化的内容反序列化到mes中。
	if err != nil {
		fmt.Println("反序列化 失败 err = ", err)
		return
	}

	return
}
