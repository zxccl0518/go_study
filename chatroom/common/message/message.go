package message

import "fmt"

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
)

type Message struct {
	Type string `json:"type"` // 消息类型
	Data string `json:"data"` //消息的类型
}

// 先定义2个消息。后面需要在增加。
type LoginMes struct {
	Userid   int    `json:"userId"`  //用户id
	UserPwd  string `json:"userPwd"` //用户密码
	UserName string `json:userName`  //用户名
}

type LoginResMes struct {
	Code  int    `json:"code"`  // 返回状态码 500表示该用户为注册，200表示登陆成功。
	Error string `json:"error"` // 返回错误信息。
}

func (m Message) PrintMessage() {
	fmt.Println("Message 结构体的 打印方法。")
}
