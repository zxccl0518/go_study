package message

import "fmt"

const (
	LoginMesType       = "loginType"
	LoginResMesType    = "loginResType"
	RigisterMesType    = "rigisterType"
	RigisterResMesType = "rigisterResType"
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

func (mes Message) Print() {
	fmt.Println("Message 结果提 print的方法。 --- ")
}

func (mes LoginResMessage) Print() {
	fmt.Println("LoginResMessage 结果提 print的方法。 --- ")
}
