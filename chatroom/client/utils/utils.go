package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

	"github.com/zxccl0518/go_study/chatroom/common/message"
)

// 这里我们将这些方法，关联到结构体中。
type Transfer struct {
	// 分析应该有哪些字段。
	Conn net.Conn   // 连接
	Buf  [8064]byte // 传输时 使用的缓冲。
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	// buf := make([]byte, 1024*8)

	// conn.Read在conn没有关闭的情况下，才会阻塞。
	// 如果客户端关闭了，conn则不会阻塞
	n, err := this.Conn.Read(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Read() failed, err =", err)
		return
	}

	// 根据buf[:4] 转成一个uint32类型
	var pkgLen = binary.BigEndian.Uint32(this.Buf[:n])
	// fmt.Println("客户端发送过来的消息 长度为 = ", int(pkgLen))
	// fmt.Println("客户端发送过来的关于长度的 消息 内容为 = ", buf[:n])

	// 根据pkgLen 读取消息内容
	n, err = this.Conn.Read(this.Buf[:int(pkgLen)])
	if err != nil || n != int(pkgLen) {
		fmt.Println("服务器接收端，接收到的消息不完整。 err = ", err)
		return
	}

	// fmt.Println("服务器接收到 登录消息 打印的是字节切片形式的内容 = ", buf[:n])
	// fmt.Println("服务器接收到 登录消息 打印的是string形式的内容 = ", string(buf[:n]))
	// 技术就是一层窗户纸。 不弄清楚的话 永远都是个谜题。弄清楚了就如同 捅破窗户纸一般 简单。
	err = json.Unmarshal(this.Buf[:n], &mes) // 这个位置一定要用 &mes 才会将序列化的内容反序列化到mes中。
	if err != nil {
		fmt.Println("client/utils.go ReadPkg() 反序列化 失败 err = ", err)
		return
	}

	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	var pkgLen = uint32(len(data))
	// var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[:4], pkgLen)
	// 发送长度
	_, err = this.Conn.Write(this.Buf[:4])
	if err != nil {
		fmt.Println("发送消息的长度失败 err  = ", err)
		return
	}

	_, err = this.Conn.Write(data)
	if err != nil {
		fmt.Println("发送 消息本身失败 err = ", err)
		return
	}

	return
}
