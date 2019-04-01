package main

import (
	"fmt"
	"io"
	"net"

	"github.com/zxccl0518/go_study/chatroom/common/message"
	"github.com/zxccl0518/go_study/chatroom/server/process"
	"github.com/zxccl0518/go_study/chatroom/server/utils"
)

type Processor struct {
	Conn net.Conn
}

//编写一个ServerProcessMes()函数
// 功能：根据客户端发送消息种类不同，决定调用哪个函数来处理。
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	// 看看是否能接收到 客户端发送的群发的消息。
	fmt.Println("server provessor mes = ", mes)

	switch mes.Type {
	case message.LoginMesType:
		// 创建 userProcess 实例。
		userProcess := &process.UserProcess{
			Conn: this.Conn,
		}

		// 处理登录的逻辑
		err = userProcess.ServerProcessLogin(mes)
		if err != nil {
			fmt.Println("处理登录失败 ... ")
		}
	case message.RigisterMesType:
		// 创建 userProcess 实例。
		userProcess := &process.UserProcess{
			Conn: this.Conn,
		}
		err = userProcess.ServerProcessRigister(mes)
		if err != nil {
			fmt.Println("用户注册 失败。 err = ", err)
		}
	default:
		// 消息类型不存在
		fmt.Println("消息类型不存在，无法处理...s")
	}

	return
}

func (this *Processor) process2() (err error) {
	// 循环的接收客户发送的消息。
	for {
		// 这里我们阿静读取数据包，直接封装成一个函数readPkg，返回Message,Err
		// 创建一个Transfer实例，完成读包任务。
		tf := utils.Transfer{
			Conn: this.Conn,
		}

		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端断开连接---")
				return err
			}

			fmt.Println("读取客户端 传过来的登陆消息失败，c错误原因非客户端掉线。 err = ", err)
			return err
		}

		err = this.serverProcessMes(&mes)
		if err != nil {
			fmt.Println(" 服务器解析 消息失败 err= ", err)
		}
	}
}
