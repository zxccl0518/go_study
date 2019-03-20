package message

import "fmt"

const (
	LoginMesType    = "loginType"
	LoginResMesType = "loginResType"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMessage struct {
	UserID   string `json: "userID"`
	UserPwd  string `json: "userPwd"`
	UserName string `json: "userName"`
}

type LoginResMessage struct {
	Code int   `json:"code"`
	Err  error `json:"err"`
}

func (mes Message) Print() {
	fmt.Println("Message 结果提 print的方法。 --- ")
}

func (mes LoginResMessage) Print() {
	fmt.Println("LoginResMessage 结果提 print的方法。 --- ")
}
