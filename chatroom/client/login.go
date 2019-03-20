package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

	"github.com/zxccl0518/go_study/chatroom/common/message"
)

// 用户登录函数
func login(userID string, userPwd string) error {
	conn, err := net.Dial("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("拨号，连接服务器失败 err = ", err)
		return err
	}

	var mes message.Message
	mes.Type = message.LoginMessageType
	var loginMes message.LoginMessage
	loginMes.UserID = userID
	loginMes.UserPwd = userPwd

	// 将loginMes 序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("序列化失败 err = ", err)
		return err
	}
	mes.Data = string(data)

	// 将Message 序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("对Message 结构体 序列化失败 err = ", err)
		return err
	}

	var pkdLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkdLen)

	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("给 服务器发送登录信息的长度失败 err = ", err)
		return err
	}

	fmt.Println("给服务器发送的登录信息的长度 = pkglen = ", pkdLen)

	n, err = conn.Write(data)
	if err != nil || n != len(data) {
		fmt.Println("发送 登录信息失败 err = ", err)
		return err
	}

	return nil
}
