package main

import "fmt"

// 定义/ 声明一个借口
type Usb interface {
	start()
	stop()
}

// 若存在这样的一个借口,那么下面的 Phone结构体和 Camera结构体 也实现了Usb2的借口。
// go中实现接口 依据的释方法，而不是指定哪个接口的哪个方法。
type Usb2 interface {
	start()
	stop()
	work() // 但是如果Usb2接口又增加了一个 work()方法，那么Phone、Camera()2个结构体 就是没有实现Usb2接口，因为没有实现Usb2接口中所有的方法。
}

//  定义 手机结构体
type Phone struct {
}

// Phone结构体 实现了 Usb接口中的所有方法。， 即为实现了 Usb接口。
func (p Phone) start() {
	fmt.Println("手机 实现接口，开始工作")
}
func (p Phone) stop() {
	fmt.Println("手机 实现接口，停止工作")
}

// 定义 相机结构体
type Camera struct {
}

// Camera结构体 实现了 Usb接口中的所有方法。， 即为实现了 Usb接口。
func (c Camera) start() {
	fmt.Println("相机 实现接口，开始工作")
}
func (c Camera) stop() {
	fmt.Println("相机 实现接口，停止工作")
}

type Computer struct {
}

func (c Computer) Working(u Usb) {
	// func (c Computer) Working(u Usb2) { // 如果 接口Usb2 定义了work()方法，但是Camera这个结构体并没有实现work()这个方法，那么这个时候就会出现错误。   Camera does not implement Usb2 (missing work method)
	u.start()
	u.stop()
}

func main() {
	// var phone = Phone{}
	// phone.start()
	// phone.stop()

	var camera = Camera{}
	// camera.start()
	// camera.stop()

	computer := Computer{}
	computer.Working(camera)
}
