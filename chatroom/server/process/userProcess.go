package process

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/zxccl0518/go_study/chatroom/common/message"
	"github.com/zxccl0518/go_study/chatroom/server/model"
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
	// 我们需要到redis数据库完成验证。
	// 使用modelMyUserDao 到redis去验证。
	_, err = model.MyUserDao.Login(loginMes.UserID, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOT_EXISTS {
			loginResMes.Code = 500
			// 这里我们先设置成功，然后我们再 根绝具体的信息返回具体的
			// loginResMes.Err = "该用户不存在，请注册在使用。"
			loginResMes.Err = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Err = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Err = "服务器内部错误。"
		}
	} else {
		loginResMes.Code = 200
		loginResMes.Err = "登录成功。"
	}
	// // 如果用户的id为100  密码为123456，我们认为合法，否则不合法。
	// if loginMes.UserID == "100" && loginMes.UserPwd == "123456" {
	// 	//合法
	// 	loginResMes.Code = 200
	// } else {
	// 	// 不合法
	// 	loginResMes.Code = 500
	// 	// loginResMes.Err = errors.New("该用户不存在，请注册在使用。 ")
	// 	loginResMes.Err = "该用户不存在，请注册在使用。"
	// }

	// 将LoginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("process包，ServerProcessLogin() --- 序列化失败。 err = ", err)
	}

	// 4.将data 赋值给resMes结构体
	resMes.Data = string(data)

	// 5。对resMes进行序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("process包，ServerProcessLogin() --- json.Marshal fail err = ", err)
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

// 注册用户。
func (this *UserProcess) ServerProcessRigister(mes *message.Message) (err error) {
	var rigisterMes message.RigisterMes

	err = json.Unmarshal([]byte(mes.Data), &rigisterMes)
	if err != nil {
		fmt.Println("process包，ServerProcessRigister() --- 注册信息 反序列化失败。 err = ", err)
		return
	}

	fmt.Println("打印 用户注册的信息 ------------------------")
	fmt.Printf("userID:%d, userPwd:%s, userName:%s\n", rigisterMes.User.UserID, rigisterMes.User.UserPwd, rigisterMes.User.UserName)

	var resMes message.Message
	resMes.Type = message.RigisterMesType
	var rigisterResMes message.RigisterResMes
	err = model.MyUserDao.Rigister(&rigisterMes)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			fmt.Println("用户已经存在，注册失败 err = ", err.Error())
			rigisterResMes.Code = 505
			rigisterResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			rigisterResMes.Code = 506
			rigisterResMes.Error = "注册发生错误。"
		}
	} else {
		rigisterResMes.Code = 200
	}

	data, err := json.Marshal(rigisterResMes)
	if err != nil {
		fmt.Println("process包，ServerProcessRigister() --- 序列化失败。")
		return
	}
	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("process包，ServerProcessRigister() --- 序列化失败 err= ", err)
		return
	}

	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	tf.WritePkg(data)

	return
}
