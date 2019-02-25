package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 1.声明一个Hero结构体
type Hero struct {
	Name string
	Age  int
}

// 2.声明一个Hero结构体切片类型。
type HeroSlice []Hero

// 3.实现Interface接口
func (hs HeroSlice) Len() int {
	return len(hs)
}

// less方法就是决定你使用 升序还是降序 作为排序标准。
// 这里我们安装Hero的年龄从小到大
func (hs HeroSlice) Less(i, j int) bool {
	// 按照年龄 的升序进行排序。
	// return hs[i].Age < hs[j].Age
	// 修改成 对名字的排序
	return hs[i].Name < hs[j].Name
}

func (hs HeroSlice) Swap(i, j int) {
	// temp := hs[i]
	// hs[i] = hs[j]
	// hs[j] = temp
	// 等价于下面这句话
	hs[i], hs[j] = hs[j], hs[i]
}

func main() {
	// 实现对Hero结构体切片的最佳实现

	// 定义一个数组/ 切片。
	var intSlice = []int{0, -1, 10, 7, 90}
	// 要求对intSlice切片进行排序。
	// 1.冒泡排序。。。
	// 2.也可以 使用系统提供的方法。		sort.Ints([]slice) 是系统提供的 对int切片类型的变量 进行排序，默认是升序
	sort.Ints(intSlice)
	fmt.Println("intSlice = ", intSlice)

	// 请大家对结构体切片进行排序。
	// 1.冒泡排序。
	// 2.也可以使用系统提供的方法。

	var hs HeroSlice
	for i := 0; i < 10; i++ {
		hs = append(hs, Hero{fmt.Sprintf("英雄<>~%d", rand.Intn(100)), rand.Intn(100)})
	}

	// 看看排序前的顺序。
	for i, v := range hs {
		fmt.Printf("hs[%d] = %v \n", i, v)
	}

	sort.Sort(hs)
	// 看看排序后的顺序。
	for _, v := range hs {
		fmt.Printf("排序后 ->  = %v \n", v)
	}

	// 面试题
	i := 10
	j := 20
	i, j = j, i
	fmt.Println("i = ", i, ", j = ", j)
}
