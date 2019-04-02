package message

import (
	"fmt"
)

const (
	LoginMesType            = "loginType"
	LoginResMesType         = "loginResType"
	RigisterMesType         = "rigisterType"
	RigisterResMesType      = "rigisterResType"
	NotifyUserStatusMesType = "NotifyUserStatusMesType"
	SmsMesType              = "SmsMesType"
	SmsPrivateChatMesType   = "SmsPrivateChatMesType"
)

// 这里我们定义几个用户的状态常量。
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMessage struct {
	UserID   int    `json:"userID"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMessage struct {
	Code    int `json:"code"`
	UsersID []int
	Err     string `json:"err"`
}

type RigisterMes struct {
	User User `json:"user"`
}

type RigisterResMes struct {
	Code  int    `json:"code"` // 返回状态码400 表示该用户已经被占有了，200表示注册成功
	Error string `json:"error"`
}

// 为了配合服务器推送用户状态变化的消息
type NotifyUserStatusMes struct {
	UserID int `json:"userID`  // 用户的id
	Status int `json:"status"` // 用户的状态。
}

// 增加一个SmsMes // 发送的消息
type SmsMes struct {
	Content string `json:"content"`
	User           // 匿名结构体。
}

type SmsPrivateChatMes struct {
	Content string
	UserMe  User
	UserYou User
}

// type SmsReMes struct {
// }

func (mes Message) Print() {
	fmt.Println("Message 结果提 print的方法。 --- ")
}

func (mes LoginResMessage) Print() {
	fmt.Println("LoginResMessage 结果提 print的方法。 --- ")
}
