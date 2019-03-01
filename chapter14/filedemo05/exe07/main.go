package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	ChCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
}

func (c *CharCount) countCh(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("打开文件 失败。 err = ", err)
		return
	}

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		charTable := []rune(str)
		for _, v := range charTable {

			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				c.ChCount++
			case v >= '0' && v <= '9':
				c.NumCount++
			case v == ' ' || v == '\t':
				c.SpaceCount++
			default:
				c.OtherCount++
			}
		}
	}

	fmt.Printf("字符的个数为%v， 数字的个数为%v， 空格的个数为：%v， 其他类型的个数为%v \n", c.ChCount, c.NumCount, c.SpaceCount, c.OtherCount)
}

func main() {
	//思路 打开一个文件，创建一个Reader
	// 每读取一行，就去统计改行有多少个 英文，数字，空格，和其他字符。
	// 然后将结果保存到一个接头体中。

	filePath := "e:/abc.txt"
	charCount := &CharCount{
		0, 0, 0, 0,
	}
	charCount.countCh(filePath)
}
