package main

import "fmt"

type ValNode struct {
	row int
	col int
	val interface{}
}

func main() {
	// 1.先创建一个原始数组。
	var cheesMap [11][11]int
	cheesMap[1][2] = 1 // 黑子
	cheesMap[2][3] = 2 // 篮子

	// 2.输出看原始数组。
	for _, v := range cheesMap {
		for _, v2 := range v {
			fmt.Print(v2, "\t")
		}
		fmt.Println()
	}

	// 3.转成稀疏数组。 想 -> 算法。
	// 思路
	// 1) 遍历chessMap数组，如果发现有一个元素的值不为0, 我们创建一个node结构体。
	var sparseArr []ValNode
	// 标准的稀疏数组 应该还有一个记录原始的二维数组的规模(行和列，默认值)
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)
	for i, v := range cheesMap {
		for j, v2 := range v {
			if v2 != 0 {
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}

	// 输出稀疏数组
	for i, valNode := range sparseArr {
		fmt.Printf("%d : %d %d %d \n", i, valNode.row, valNode.col, valNode.val)
	}
}
