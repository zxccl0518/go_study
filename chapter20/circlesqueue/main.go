package main

import (
	"errors"
	"fmt"
	"os"
)

// 判断 环形队列是不是满了: (tail+1)%maxSize == head
// 判断 环形队列是不是空的  tail == head
// 判断 环形队列 元素个数。 (tail+maxSize-1)%maxSize
type CircleSqueue struct {
	maxSize int
	head    int
	tail    int
	array   [5]int
}

// 队列 入列
func (this *CircleSqueue) Push(val int) (err error) {
	result := this.isFull()
	if !result {
		this.array[this.tail] = val
		this.tail = (this.tail + 1) % this.maxSize
		return
	} else {
		return errors.New("queue is full")
	}
}

// 队列 出列
func (this *CircleSqueue) Pop() (val int, err error) {
	if !this.isEmpty() {
		val = this.array[this.head]
		this.head = (this.head + 1) % this.maxSize

		return
	}
	return 0, errors.New("队列 已经空了")
}

// 显示队列。
func (this *CircleSqueue) ListQueue() {
	tempHead := this.head

	if !this.isEmpty() {
		for i := 0; i < this.Size(); i++ {
			fmt.Println(this.array[tempHead])
			tempHead = (tempHead + 1) % this.maxSize
		}
	} else {
		fmt.Println("队列已经为空， 不做显示。")
	}
}

// 判断环形队列是否满了
func (this *CircleSqueue) isFull() (res bool) {
	res = false
	if (this.tail+1)%this.maxSize == this.head {
		res = true
	}

	return
}

// 判断环形队列是否空了
func (this *CircleSqueue) isEmpty() (res bool) {
	res = false
	if this.head == this.tail {
		res = true
	}

	return
}

func (this *CircleSqueue) Size() int {
	return (this.tail + this.maxSize - this.head) % this.maxSize
}

func main() {
	// 初始化一个 环形队列。
	circleQueue := &CircleSqueue{
		head:    0,
		tail:    0,
		maxSize: 5,
	}

	var key string
	var val int
	for {
		fmt.Println("1.请输入add 表示添加数据到队列。")
		fmt.Println("2.请输入get 表示添加数据到队列。")
		fmt.Println("3.请输入show 表示添加数据到队列。")
		fmt.Println("4.请输入exit 表示添加数据到队列。")
		fmt.Scanf("%s\n", &key)
		switch key {
		case "add":
			fmt.Println("请输入 数字，添加到队列。")
			fmt.Scanf("%d\n", &val)
			err := circleQueue.Push(val)
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println("添加数据到队列成功。")
		case "get":
			res, err := circleQueue.Pop()
			if err != nil {
				fmt.Println("pop 出队列失败， err= ", err)
			} else {
				fmt.Println("出列元素为 = ", res)
			}
		case "show":
			fmt.Println("-----------------------------------")
			circleQueue.ListQueue()
			fmt.Println("-----------------------------------")
		case "exit":
			os.Exit(0)
		default:
		}
	}
}
