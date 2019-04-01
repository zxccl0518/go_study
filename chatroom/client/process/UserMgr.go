package process

import (
	"fmt"

	"github.com/zxccl0518/go_study/chatroom/client/model"

	"github.com/zxccl0518/go_study/chatroom/common/message"
)

type UserMgr struct {
}

var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser // 我们在用户登录成功后，完成对CurUser初始化

// 编写一个方法，处理返回的NotifyUserStatusMes
func updateUserstatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	// 适当优化
	user, ok := onlineUsers[notifyUserStatusMes.UserID]
	if !ok { // 原来没有
		user = &message.User{
			UserID: notifyUserStatusMes.UserID,
		}
	}

	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserID] = user
}

// 在客户端显示当前在线的用户
func outputOnlineUser() {
	fmt.Println("当前在线用户列表 服务器 主动通知客户端 ------------------------")
	for id, user := range onlineUsers {
		if user.UserStatus == message.UserOnline {
			fmt.Println("用户id = \t message.UserOnline = ", id, user.UserStatus)
		}
	}
}
