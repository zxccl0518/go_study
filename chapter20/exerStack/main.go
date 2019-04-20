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

// 判断一个字符是不是运算符【+ - * /】。
func (this *Stack) isOper(value int) bool {
	if value == 42 || value == 43 || value == 45 || value == 47 {
		return true
	}

	return false
}

// 判断 运算符的优先级
// [* /  >=  + -]
func (this *Stack) Priority(oper int) int {
	res := 0
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 0
	} else {
		fmt.Println("其他 运算符情况 不做计算。")
	}

	return res
}

// 运算方法。
func (this *Stack) Cal(num1 int, num2 int, oper int) int {
	result := 0
	switch oper {
	case 42:
		result = num2 * num1
	case 43:
		result = num2 + num1
	case 45:
		result = num2 - num1
	case 47:
		result = num2 / num1
	default:
		fmt.Println("运算符错误。")
	}

	return result
}

func main() {
	numberStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}

	operatorStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}

	// 表达式。
	exp := "3+2*6-2"
	// 定义一个index 帮助我们扫描exp。
	index := 0
	// 为了配合运算，我们定义需要的变量。
	num1 := 0
	num2 := 0
	oper := 0
	result := 0
	for {
		ch := exp[index : index+1]            // ch是一个字符串。只不过只有一个字符的字符串。
		temp := int([]byte(ch)[0])            // 就是字符串对应的ASCII码
		if numberStack.isOper(temp) == true { // 说明是符号。
			// 这里有2个逻辑
			// 1 如果operStack是一个空栈。直接入栈
			if operatorStack.Top == -1 { // 空栈。
				operatorStack.Push(temp)
			} else {
				if operatorStack.Priority(operatorStack.arr[operatorStack.Top]) >= operatorStack.arr[temp] {
					num1, _ = numberStack.Pop()
					num2, _ = numberStack.Pop()
					oper, _ = operatorStack.Pop()
					result = operatorStack.Cal(num1, num2, oper)
					numberStack.Push(result)
				} else {
					operatorStack.Push(temp)
				}
			}
		} else { // 说明是数字
			numberStack.Push(temp)
		}

		// 判断 index 是都已经扫描到了计算表达式的最后了。
		if index+1 >= len(exp) {
			break
		}
		// 继续扫描。
		index++
	}

	for {
		if operatorStack.arr[operatorStack.Top] == -1 {
			break
		}
		num1, _ = numberStack.Pop()
		num2, _ = numberStack.Pop()
		oper, _ = operatorStack.Pop()
		result = operatorStack.Cal(num1, num2, oper)
		numberStack.Push(result)
	}

	result, _ = numberStack.Pop()

	fmt.Println("运算结果 为 result  = ", result)
}
