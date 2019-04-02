package process

import (
	"encoding/json"
	"fmt"

	"github.com/zxccl0518/go_study/chatroom/common/message"
)

type SmsMgr struct {
}

func (this *SmsMgr) outputGroupMes(mes *message.Message) (err error) {
	var smsMes message.SmsMes

	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("SmsMgr outputGroupMes() 反序列化消息失败 err = ", err)
		return
	}

	content := fmt.Sprintf("来自用户id[%v] 对大家说：%v ", smsMes.UserID, smsMes.Content)
	fmt.Println(content)

	return
}

func (this *SmsMgr) OutPutPrivateChatMes(mes *message.Message) (err error) {
	var smsPrivateMes message.SmsPrivateChatMes
	err = json.Unmarshal([]byte(mes.Data), &smsPrivateMes)
	if err != nil {
		fmt.Println("smsMgr.go  OutPutPrivateChatMes() 反序列化失败 err = ", err)
		return
	}

	var content = fmt.Sprintf("来自用户[%v] 对你说：%v", smsPrivateMes.UserMe.UserID, smsPrivateMes.Content)
	fmt.Println(content)
	fmt.Println()

	return
}
