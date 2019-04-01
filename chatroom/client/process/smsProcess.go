package process

import (
	"encoding/json"
	"fmt"

	"github.com/zxccl0518/go_study/chatroom/client/utils"

	"github.com/zxccl0518/go_study/chatroom/common/message"
)

type SmsProcess struct {
}

func (this *SmsProcess) sendGroups(content string) (err error) {
	//1. 创建一个message
	var mes message.Message
	mes.Type = message.SmsMesType

	// 2.创建一个smsMes
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserID = CurUser.UserID
	smsMes.UserName = CurUser.UserName
	fmt.Println("=======>>> smsMes.Content = ", smsMes.Content)
	fmt.Println("=======>>> smsMes.UserName = ", smsMes.UserName)

	//3. 序列化smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("sendGroundMes json.Marshal fail err = ", err.Error())
		return
	}

	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("sendGroundMes json.Marshal fail err = ", err.Error())
		return
	}

	tf := utils.Transfer{
		Conn: CurUser.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("sendGroupMes writePkg err = ", err)
		return
	}

	return
}
