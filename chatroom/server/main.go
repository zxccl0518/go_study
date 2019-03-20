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
	// fmt.Println("客户端发送过来的消息 长度为 = ", int(pkgLen))
	// fmt.Println("客户端发送过来的关于长度的 消息 内容为 = ", buf[:n])

	// 根据pkgLen 读取消息内容
	n, err = conn.Read(buf[:int(pkgLen)])
	if err != nil || n != int(pkgLen) {
		fmt.Println("服务器接收端，接收到的消息不完整。 err = ", err)
		return
	}

	// fmt.Println("服务器接收到 登录消息 打印的是字节切片形式的内容 = ", buf[:n])
	// fmt.Println("服务器接收到 登录消息 打印的是string形式的内容 = ", string(buf[:n]))
	// 技术就是一层窗户纸。 不弄清楚的话 永远都是个谜题。弄清楚了就如同 捅破窗户纸一般 简单。
	err = json.Unmarshal(buf[:n], &mes) // 这个位置一定要用 &mes 才会将序列化的内容反序列化到mes中。
	if err != nil {
		fmt.Println("反序列化 失败 err = ", err)
		return
	}

	return
}

func writePkg(conn net.Conn, data []byte) (err error) {
	// var loginResMes message.LoginResMessage
	// loginResMes.Code = 200
	// loginResMes.Err = nil

	// data, err := json.Marshal(loginResMes)
	// if err != nil {
	// 	fmt.Println("序列化 失败 err = ", err)
	// 	return
	// }

	// _, err = conn.Write(data)
	// if err != nil {
	// 	fmt.Println("服务器 返回给客户端 消息失败。。。")
	// 	return
	// }

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

// 编写一个 ServerProcessLogin函数，专门处理登录的请求
func ServerProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	// 核心代码
	// 1.先从mes中取出mes.Data,并反序列化成LoginMes
	var loginMes message.LoginMessage
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Umnmarshal fail err = ", err)
		return
	}

	//1. 先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	// 2. 再声明一个 LoginResMes
	var loginResMes message.LoginResMessage

	// 如果用户的id为100  密码为123456，我们认为合法，否则不合法。
	if loginMes.UserID == "100" && loginMes.UserPwd == "123456" {
		//合法
		loginResMes.Code = 200
	} else {
		// 不合法
		loginResMes.Code = 500
		// loginResMes.Err = errors.New("该用户不存在，请注册在使用。 ")
		loginResMes.Err = "该用户不存在，请注册在使用。 "
	}

	// 将LoginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("序列化失败。 err = ", err)
	}

	// 4.将data 赋值给resMes结构体
	resMes.Data = string(data)

	// 5。对resMes进行序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail err = ", err)
		return
	}

	// 发送data 我们将其功能封装到writePkg函数中。
	err = writePkg(conn, data)
	return
}

//编写一个ServerProcessMes()函数
// 功能：根据客户端发送消息种类不同，决定调用哪个函数来处理。
func ServerProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		// 处理登录的逻辑
		err = ServerProcessLogin(conn, mes)
		if err != nil {
			fmt.Println("处理登录失败=== ")
		}
	case message.RigisterMesType:

	default:

	}

	return
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

		err = ServerProcessMes(conn, &mes)
		if err != nil {
			fmt.Println(" 服务器解析 消息失败 err= ", err)
		}
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
