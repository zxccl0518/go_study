package process

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/zxccl0518/go_study/chatroom/server/utils"

	"github.com/zxccl0518/go_study/chatroom/common/message"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMes(mes *message.Message) (err error) {
	var smsMes message.SmsMes
	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Printf("sendGroupMes Unmashal failed err = %v\n", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Printf("sendGroupMes Marshal failed err = %v\n", err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		if id == smsMes.UserID {
			continue
		}

		// 将消息发送给其他在线好友
		this.sendMesToUserOnline(data, up.Conn)
	}

	return
}

func (this *SmsProcess) sendMesToUserOnline(data []byte, conn net.Conn) (err error) {
	tf := &utils.Transfer{
		Conn: conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Printf("sendMesToEachOnline writePkg failed err  = %v\n", err)
		return
	}

	return
}

func (this *SmsProcess) SendPrivateMes(mes *message.Message) (err error) {
	// 提取 smsPrivateMes 取出私聊对象的userID，并判断对方是否在线。
	var smsPrivateMes message.SmsPrivateChatMes
	err = json.Unmarshal([]byte(mes.Data), &smsPrivateMes)
	if err != nil {
		fmt.Println("json.Unmarshal() failed err = ", err)
		return
	}

	if up, ok := userMgr.onlineUsers[smsPrivateMes.UserYou.UserID]; ok {
		data, err := json.Marshal(mes)
		if err != nil {
			fmt.Println("smsProcess.go SendPrivateMes() 反序列化mes 失败,err = ", err)
			// return
		} else {
			err = this.sendMesToUserOnline(data, up.Conn)
		}
	} else {
		fmt.Println("smsProcess.go sendPrivateMes() 私聊对象没有在线。")
	}

	return
}
