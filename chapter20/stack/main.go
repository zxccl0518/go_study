package main

import (
	"errors"
	"fmt"
)

// 使用数组来模拟一个栈的使用。
type Stack struct {
	MaxTop int    // 表示我们栈最大都可以存储的个数。
	Top    int    // 表示栈顶，因为我们栈顶固定，因此我们使用top
	arr    [5]int // 数组模拟栈
}

// 入栈操作。
func (this *Stack) Push(value int) (err error) {
	if this.Top >= this.MaxTop-1 {
		fmt.Println("栈 已经满了，无法再入栈。")
		return
	}
	this.Top++
	this.arr[this.Top] = value

	return
}

// 出栈操作
func (this *Stack) Pop() (value int, err error) {
	if this.Top == -1 {
		fmt.Println("栈已经为空，无法继续出栈")
		return -1, errors.New("栈已经为空，无法继续出栈")
	}

	value = this.arr[this.Top]
	this.Top--

	return value, nil
}

func (this *Stack) List() {
	if this.Top == -1 {
		fmt.Println("栈 为空，无法打印。")
		return
	}

	for i := this.Top; i >= 0; i-- {
		fmt.Printf("this.top = %d, value = %v\n", i, this.arr[i])
	}
}

func main() {
	stack := &Stack{
		Top:    -1,
		MaxTop: 5,
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	stack.List()

	stack.Pop()
	stack.Pop()

	fmt.Println("出栈操作2次 之后打印的值")
	stack.List()
}
