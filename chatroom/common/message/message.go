package message

const (
	LoginMessageType    = "loginType"
	LoginResMessageType = "loginResType"
)

type Message struct {
	Type string
	Data string
}

type LoginMessage struct {
	UserID   string
	UserPwd  string
	UserName string
}

type ResMessage struct {
	Code int
	err  error
}
