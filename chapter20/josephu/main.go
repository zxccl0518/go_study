package main

import "fmt"

type boy struct {
	No   int
	Next *boy
}

func AddBoy(n int) (head *boy) {
	head = &boy{}
	tail := head

	if n < 1 {
		return
	}

	for i := 1; i <= n; i++ {
		boy := &boy{
			No: i,
		}

		if i == 1 {
			head = boy
			head.Next = boy
			tail = head
		} else {
			tail.Next = boy
			boy.Next = head
			tail = boy
		}
	}

	return
}

func ShowBoy(head *boy) {
	if head.Next == nil {
		fmt.Println("ShowBoy() --- 链表头 指针为空。--- ")
		return
	}

	temp := head
	for {
		fmt.Println("boy:", temp.No)

		if temp.Next == head {
			break
		}

		temp = temp.Next
	}
}

// 约瑟夫 玩游戏问题。
func PlayGame(head *boy, startIndex int, countNum int) {
	if head.Next == nil {
		fmt.Println("链表头为空指针，无法进行约瑟夫游戏。")
		return
	}

	tail := head
	queueHead := head

	// 找到队尾的结点。
	for {
		tail = tail.Next
		if tail.Next == head {
			break
		}
	}
	fmt.Println("tail 是最后一个结点， no = ", tail.No)

	// 将queueHead指针指到要删除的按个位置上。同时尾指针也一并移动同样多的距离(也就是说保持在queueHead的后面)
	for i := 1; i <= startIndex-1; i++ {
		queueHead = queueHead.Next
		tail = tail.Next
	}

	for {
		for i := 1; i <= countNum-1; i++ {
			queueHead = queueHead.Next
			tail = tail.Next
		}

		fmt.Printf("queueHead.No = %d, tail.No = %d\n", queueHead.No, tail.No)

		queueHead = queueHead.Next
		tail.Next = queueHead

		// 说明此时此刻 仅剩下最后一个结点了。就不用循环了。
		if queueHead == tail {
			break
		}
	}
	fmt.Println("最后一个 出列的结点 no = ", queueHead.No)
}

func main() {
	boyHead := AddBoy(5)

	ShowBoy(boyHead)

	PlayGame(boyHead, 2, 2)
}
