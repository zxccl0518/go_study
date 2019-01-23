package main

import "fmt"

// 求出斐波那契数。
func test1(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return test1(n-1) + test1(n-2)
}

// test2 求函数的值。
// 已知 f(1) = 3;fn(n) = 2*f(n-1)+1
// 请用递归的思想变成，求出 f(n)
func test2(n int) int {
	if n == 1 {
		return 3
	}

	return 2*test2(n-1) + 1
}

// 练习题木3， 猴子吃桃子的问题。  一堆桃子，猴子第一天吃了其中了的一半，并且再多吃了一个！ 以后的每天猴子都吃掉了其中的一半，并且再多吃一个。当吃到帝释天的时候，想在吃的时候(还没吃)，发现只有一个桃子了。请问最初有多少个桃子。
// 思路：第十天 ：1个， 第九天:(第十天+1）*2   第八天:(第九天+1)*2 ...
func test3(n int) int {

	if n == 10 {
		return 1
	} else {
		return (test3(n+1) + 1) * 2
	}
}

func main() {
	// 递归的方式，给以个整数n，求出它的 斐波那契数 1,1,2,3,5,8,13
	// 思路： 1) 当 n == 1 || n == 2的时候，return 1
	//       2)

	// fmt.Println("sum = ", test1(4))
	// fmt.Println("sum = ", test1(5))
	// fmt.Println("sum = ", test1(6))
	// fmt.Println("sum = ", test1(7))

	//---------------------------------

	// fmt.Println("test2 sum = ", test2(2))
	// fmt.Println("test2 sum = ", test2(3))
	// fmt.Println("test2 sum = ", test2(4))
	// fmt.Println("test2 sum = ", test2(5))

	// 求第十天的桃子的数量。
	fmt.Println("test3 第十天，sum = ", test3(10))
	fmt.Println("test3 第九天，sum = ", test3(9))
	fmt.Println("test3 第八天，sum = ", test3(8))
	fmt.Println("test3 第1天，sum = ", test3(1))

}
