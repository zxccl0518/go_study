package process

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/zxccl0518/go_study/chatroom/common/message"

	"github.com/zxccl0518/go_study/chatroom/client/utils"
)

// type

// 显示登录成功后的界面。
func ShowMenu() {
	fmt.Println("--------- 恭喜 xxx 登录成功 --------")
	fmt.Println("\t\t\t 1.显示在线用户列表\t\t\t")
	fmt.Println("\t\t\t 2.发送消息。\t\t\t")
	fmt.Println("\t\t\t 3.信息列表。\t\t\t")
	fmt.Println("\t\t\t 4.私人聊天。\t\t\t")
	fmt.Println("\t\t\t 5.退出系统。\t\t\t")
	fmt.Println("\t\t\t 请选择(1-4):")
	var key int
	// var content string

	// 因为我们总会使用到SmsProcess实例，因此我们将其定义在switch外部.
	smsProcess := &SmsProcess{}
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("显示 在线用户列表")
		outputOnlineUser()
	case 2:
		fmt.Println("你想对大家说点什么:")
		reader := bufio.NewReader(os.Stdin)
		data, _, _ := reader.ReadLine()
		// fmt.Scanln(&content)

		// 这种方式，可以解决输入字符串中带空格的问题。
		smsProcess.sendGroups(string(data))
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("私聊系统，与个人单独聊天")
		fmt.Print("请输入 私聊用户的 用户id:")
		var userID int
		var flag = false
		fmt.Scanf("%d\n", &userID)
		for id := range onlineUsers {
			if id == userID {
				flag = true
				break
			}
		}

		if flag == true {
			fmt.Print("请输入 要发送的内容:")
			newreader := bufio.NewReader(os.Stdin)
			content, _, _ := newreader.ReadLine()

			smsProcess := &SmsProcess{}
			err := smsProcess.sendPrivateChat(string(content), userID)
			if err != nil {
				fmt.Println("发送私聊失败。 err = ", err)
			}
		}
	case 5:
		fmt.Println("你 退出了系统")
		os.Exit(0)
	default:
	}
}

// 和服务器端保持通讯.
func serverProcessMes(conn net.Conn) {
	// 创建一个 Transfer实例。不停的读取 服务器发送消息。
	// 1.只要服务器 和 客户端不断开 连接。就会阻塞的方式读取。
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		// 客户端不停地 读取
		fmt.Println("客户端正在等待读取服务器 发送的消息。")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err = ", err)
			return
		}

		// 如果读取到了，又是下一步处理逻辑。
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			// 有人上线了
			//1. 取出NotifyUserStatusMes
			//2. 把这个用户的信息，状态 保存到客户端的Map   map[int]User中

			// 将 消息反序列化
			var notifyUsreStatusMes message.NotifyUserStatusMes
			err = json.Unmarshal([]byte(mes.Data), &notifyUsreStatusMes)
			if err != nil {
				fmt.Println("反序列化失败")
				return
			}
			// 更新 onlineUsers map
			updateUserstatus(&notifyUsreStatusMes)

			// 显示当前所有在线的好友
			outputOnlineUser()

		case message.SmsMesType:
			// var smsMes message.SmsMes
			// err = json.Unmarshal([]byte(mes.Data), &smsMes)
			// if err != nil {
			// 	fmt.Println("接受服务器端转发的群聊消息失败 --- 反序列化失败 err = ", err)
			// 	return
			// }

			// content := fmt.Sprintf("用户id:%d 跟大家说 %v", smsMes.UserID, smsMes.Content)
			// fmt.Println(content)s
			smsMgr := &SmsMgr{}
			smsMgr.outputGroupMes(&mes)
		case message.SmsPrivateChatMesType:
			fmt.Printf("客户端 用户id[%v]接收到 私聊 data = ", mes.Data)
			smg := &SmsMgr{}
			smg.OutPutPrivateChatMes(&mes)
		default:
			fmt.Printf("服务器端 返回了未知的消息类型。")
		}
	}
}
