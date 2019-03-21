package process

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/zxccl0518/go_study/chatroom/common/message"
	"github.com/zxccl0518/go_study/chatroom/server/utils"
)

type UserProcess struct {
	Conn net.Conn
}

// 编写一个 ServerProcessLogin函数，专门处理登录的请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
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
		loginResMes.Err = "该用户不存在，请注册在使用。"
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

	// 6.发送data 我们将其功能封装到writePkg函数中。
	// 因为我们使用了分层模式(mvc), 我们先创建爱你一个Transfer实例。然后来了读取。
	tf := &utils.Transfer{
		Conn: this.Conn,
	}

	err = tf.WritePkg(data)
	return
}
