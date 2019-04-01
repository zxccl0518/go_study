package model

import (
	"net"

	"github.com/zxccl0518/go_study/chatroom/common/message"
)

// 因为在客户端，很多地方会使用到curUser，我们将其作为一个全局变量。
type CurUser struct {
	Conn         net.Conn
	message.User // 匿名结构体的方式。
}
