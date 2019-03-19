package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

	message "github.com/zxccl0518/go_study/chatroom/common/message"
)

// 写一个函数，完成登录。
func login(userID int, userPwd string) (err error) {
	// 1.连接到服务器端
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	}
	// 演示关闭
	defer conn.Close()

	// 2.准备通过conn发送消息给服务器.
	var mes message.Message
	mes.Type = message.LoginMesType
	// 3.创建一个LoginMes结构体
	var loginMes message.LoginMes
	loginMes.Userid = userID
	loginMes.UserPwd = userPwd

	// 4.将loginMes序列化。
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal loginMes err = ", err)
		return
	}

	// 5.把data赋值给mes.Data字段
	mes.Data = string(data)

	// 6.将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal mes err = ", err)
		return
	}

	// 7. 到这个时候data就是我们要发送的消息。
	// 7.1 先把data 的长度，发送给服务器。
	// 现货渠道 data的长度->
	var pkgLen = uint32(len(data))
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[0:4], pkgLen)

	// 发送长度
	n, err := conn.Write(bytes[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write() err = ", err)
		return
	}
	fmt.Println("客户端 发送消息的长度成功。 ")

	// n, err = conn.Write(data)
	// if err != nil {
	// 	fmt.Println("客户端 发送登录内容失败 err = ", err)
	// 	return
	// }

	return
}
