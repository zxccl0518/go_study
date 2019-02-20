package main

import (
	"fmt"
	"sort"
)

type Student struct {
	ID   int
	NAME string
}

func main() {
	// map的声明和注意事项。
	var a map[string]string  // 仅仅定义了map，不会自动给map分配空间
	fmt.Println("map = ", a) //map =  map[]
	// map 一定要make分配空间之后才可以使用。

	a = make(map[string]string, 10)
	a["no.1"] = "宋江"
	a["no.2"] = "吴用"
	a["no.1"] = "武松"
	a["no.3"] = "吴用"
	a["no.4"] = "吴用"
	a["no.5"] = "吴用"
	a["no.6"] = "吴用"
	a["no.7"] = "吴用"
	a["no.8"] = "吴用"
	a["no.9"] = "吴用"
	a["no.10"] = "吴用"
	fmt.Println("map = ", a) //map =  map[no.1:武松 no.2:吴用]	说明，k是不可以重复的，但是值是可以重复的。
	// golang 中的map 是无顺序的。

	// map 还可以直接赋值使用。
	var city = map[string]string{"中国的no.1 城市是": "北京"}
	fmt.Println(city)

	fmt.Println("----------------------------delete--------------------------------")
	// delete(map, key) 删除指定key的 键值对。，当key存在的时候，删除此key对应的键值对。若没有该key 则不操作，不报错。
	delete(a, "no.10")
	fmt.Println("map = ", a)
	// 若是要 删除map里面的所有键值对，有2种方式，
	// 第一种就是for循环遍历所有key 逐一删除。
	// 第二个方式就是 从新make一个新的map然后将原来的map引用 指向新make出来的 map上，原来的那个map所占用的空间就变成了垃圾，会被系统自动gc 回收掉。
	// 推荐使用第二种方式。简单 方便 快捷。
	fmt.Println("-------------------------------over--------------------------------")

	var dict = make(map[int]*Student)
	dict[0] = &Student{
		ID:   1,
		NAME: "name1",
	}

	dict[1] = &Student{
		ID:   2,
		NAME: "name2",
	}

	fmt.Println("dict[0]", dict[0])
	dict[0].ID = 1001
	fmt.Println("dict[0]", dict[0])
	fmt.Println("dict:", dict)

	// map是一个引用，如果有另一个引用 引用了map 的地址，那么，另一个引用发起的改动， 也会影响map的值。
	s := dict[0]
	s.ID = 9999
	fmt.Println("dict[0]", dict[0]) // s引用变量的改动 影响了map的值。

	// 判断map 键值 是否存在。
	if v, ok := dict[0]; ok {
		fmt.Println("存在 key = 0, value = ", v)
	} else {
		fmt.Println("没有找到 key = 0, value = ")
	}

	if v, ok := dict[2]; ok {
		fmt.Println("存在 key = 2, value = ", v)
	} else {
		fmt.Println("没有找到 key = 2 ")
	}

	// map的排序。
	var ProgramLanguage = map[string]int{
		"java":             0,
		"C":                1,
		"C++":              2,
		"Python":           3,
		"C#":               4,
		"PHP":              5,
		"JavaScript":       6,
		"Visual Basic.NET": 7,
		"Perl":             8,
	}
	// 创建切片，将map的key 也就是语言存储到切片中，然后用sort.Strings()方法对key进行排序
	var SortString []string
	for k := range ProgramLanguage {
		SortString = append(SortString, k)
	}

	sort.Strings(SortString) // 会根据字幕库的顺序排序

	for _, k := range SortString {
		fmt.Printf("key = %v , value = %v\n", k, ProgramLanguage[k])
	}

	// map的嵌套使用。
	Province := make(map[string]map[string][]string) // 定义省的字典。
	City := make(map[string][]string)                // 定义市区的字典。
	Scenery := make(map[string][]string)             // 定义景区的字典

	Scenery["西安"] = []string{"秦始皇兵马俑", "大雁塔", "大唐芙蓉园"}
	Scenery["安康"] = []string{"中坝大峡谷", "香溪洞", "南宫山", "天书峡景区"}
	City["西安市区"] = []string{"新城区", "碑林区", "莲湖区", "未央区"}
	City["安康市区"] = []string{"汉滨区", "汉阴县", "石泉县", "紫阳县"}
	Province["陕西"] = City
	City["西安景区"] = Scenery["西安"]
	City["安康景区"] = Scenery["安康"]

	for k, v := range Province {
		fmt.Println(k, v)
		for x, y := range v {
			fmt.Println(x, y)
		}
	}

	// 自己练习 判断字典的键值是不是存在。
	if k, ok := Province["北京"]; ok {
		fmt.Println("北京的这个key值 是存在的 k = ", k)
	} else {
		fmt.Println("北京的这个key值 是不存在的")
	}

	if v, ok := Province["陕西"]; ok {
		fmt.Println("陕西的这个key值 是存在的 v = ", v)
	} else {
		fmt.Println("陕西的这个key值 是不存在的")
	}

	// 拓展知识：
	// 1.make 只能为slice、map、channel 3中类型 分配内存并初始化，同时返回一个有初始化值的slice、map、channel类型的引用。它并不是指针。
	// 2.new 是内建函数，用来分配内存，它的第一个参数是一个类型，不是一个值，它的返回值是一个指向新分配类型的指针。
}
