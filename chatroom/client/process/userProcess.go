package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

	"github.com/zxccl0518/go_study/chatroom/client/utils"

	"github.com/zxccl0518/go_study/chatroom/common/message"
)

type UserProcess struct {
	// 暂时不需要字段。
}

// 给关联一个用户登录的方法。
// 用户登录函数
func (this *UserProcess) Login(userID string, userPwd string) (err error) {
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
	// fmt.Println("给服务器发送的登录信息的长度 = pkglen = ", pkdLen)
	// fmt.Println("将发送给服务器的登录消息序列化后的string = ", string(data))
	n, err = conn.Write(data)
	if err != nil || n != int(pkdLen) {
		fmt.Println("发送 登录信息失败 err = ", err)
		return err
	}

	// 创建一个Transfer实例
	tf := &utils.Transfer{}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println(" 读取数据cuowu ")
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
