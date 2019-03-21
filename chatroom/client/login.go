package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

	"github.com/zxccl0518/go_study/chatroom/common/message"
)

// 用户登录函数
func login(userID int, userPwd string) (err error) {
	conn, err := net.Dial("tcp", "0.0.0.0:8887")
	if err != nil {
		fmt.Println("拨号，连接服务器失败 err = ", err)
		return err
	}

	var mes message.Message
	mes.Type = message.LoginMesType
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

	// 发送登录消息 本身
	fmt.Println("给服务器发送的登录信息的长度 = pkglen = ", pkdLen)
	fmt.Println("将发送给服务器的登录消息序列化后的string = ", string(data))
	n, err = conn.Write(data)
	if err != nil || n != int(pkdLen) {
		fmt.Println("发送 登录信息失败 err = ", err)
		return err
	}

	// // 接收 服务器端发回来的 对于登录信息的教研结果。
	// var resData = make([]byte, 1024)
	// n, err = conn.Read(resData)
	// if err != nil {
	// 	fmt.Println("读取 服务器返回的数据失败 err = ", err)
	// 	return err
	// }
	// var loginResMes message.LoginResMessage
	// err = json.Unmarshal(resData[:n], &loginResMes)
	// if err != nil {
	// 	fmt.Println("反序列化 服务器返回的消息 shibai  err = ", err)
	// 	return err
	// }
	// loginResMes.Print()

	// if loginResMes.Code == 200 {
	// 	return errors.New("200")
	// }

	mes, err = readPkg(conn)
	if err != nil {
		fmt.Println(" 读取数据cuowu ")
		return
	}

	// 将mes的data 部分反序列化成LoginResMes
	var loginResMes message.LoginResMessage
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登录成功 ")
	} else {
		fmt.Println("登录失败 loginResMes.Error = ", loginResMes.Err)
	}

	return
}
