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
		this.sendMesToEachOnline(data, up.Conn)
	}

	return
}

func (this *SmsProcess) sendMesToEachOnline(data []byte, conn net.Conn) (err error) {
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
