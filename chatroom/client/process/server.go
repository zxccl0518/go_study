package process

import (
	"fmt"
	"net"
	"os"

	"github.com/zxccl0518/go_study/chatroom/client/utils"
)

// type

// 显示登录成功后的界面。
func ShowMenu() {
	fmt.Println("--------- 恭喜 xxx 登录成功 --------")
	fmt.Println("\t\t\t 1.显示在线用户列表\t\t\t")
	fmt.Println("\t\t\t 2.发送消息。\t\t\t")
	fmt.Println("\t\t\t 3.信息列表。\t\t\t")
	fmt.Println("\t\t\t 4.提出系统。\t\t\t")
	fmt.Println("\t\t\t 请选择(1-4):")
	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("显示 在线用户列表")
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("你 退出了体统")
		os.Exit(0)
	default:
	}
<<<<<<< HEAD
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
		fmt.Printf("mes = %s \n", mes)
	}
=======

>>>>>>> dfb43143a42c3be6a91cf6afdfedde19ce523a30
}
