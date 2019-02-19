package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//1) 创建一个byte类型的26个元素的数组，分别放置'A' - 'Z'
	// 使用for循环访问所有的元素并打印出来，提示：字符数据原酸 'A'+1 -> 'B'

	// var myChar [26]byte
	// for i := 0; i < 26; i++ {
	// 	myChar[i] = 'A' + byte(i) // 需要将 i(int) => byte
	// }

	// for _, v := range myChar {
	// 	fmt.Printf("%c、", v)
	// }

	// 请求出一个数组的最大值，并得到最大的下标
	// 假定第一个是最大值，然后从第二个元开始循环比较，如果发现有更大的，则交换。
	// var intArr [5]int = [5]int{1, -1, 9, 90, 11}
	// maxVal := intArr[0]
	// maxValIndex := 0
	// for i := 0; i < len(intArr); i++ {
	// 	if maxVal < intArr[i] {
	// 		maxVal = intArr[i]
	// 		maxValIndex = i
	// 	}
	// }
	// fmt.Printf("maxVal = %v, maxValIndex = %v \n", maxVal, maxValIndex) //maxVal = 90, maxValIndex = 3

	// 请求出一个数组的和 、平均值。 for - range
	// 思路： 求和、求平均值。
	var intArr2 = [...]int{1, -2, 9, 90, 11}
	sum := 0
	for i := 0; i < len(intArr2); i++ {
		sum += intArr2[i]
	}
	avarageValue := float64(sum) / float64(len(intArr2))
	fmt.Printf("sum = %v, avarageValue = %v\n", sum, avarageValue)

	// 随机生成5个数，并将其翻转打印。
	// rand.Intn()
	// 当我们得到随机数后，就放到一个数组中。 int数组。
	// 翻转打印。交换的次数是 len/2,倒数第一个和第一个元素交换，倒数第二个和第二个元素交换。 以此类推。。

	// 为了每次，随机出来的随机数不一致，给随机数喂一个时间的种子。
	rand.Seed(time.Now().UnixNano())
	var intArr3 [5]int
	for i := 0; i < len(intArr3); i++ {
		intArr3[i] = rand.Intn(100)
	}
	fmt.Println("交换前，intArr3 = ", intArr3)

	// 将数组位置调整，倒序排列
	temp := 0 // 作为临时变量。
	intArrLen := len(intArr3)
	for i := 0; i < intArrLen/2; i++ {
		temp = intArr3[intArrLen-1-i]
		intArr3[intArrLen-1-i] = intArr3[i]
		intArr3[i] = temp
	}
	fmt.Println("交换后，intArr3 = ", intArr3)
}
