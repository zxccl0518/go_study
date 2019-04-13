package main

import "fmt"

type HeroNode struct {
	no       int
	name     string
	nickname string
	next     *HeroNode // 表示指向下一个结点
	pre      *HeroNode // 表示指向前一个结点
}

// 插入方式，默认插的链表的尾部。
func InsertHeroNode(head *HeroNode, newNode *HeroNode) {
	temp := head
	if temp == nil {
		fmt.Println("该 链表为空，无法进行插入操作。")
		return
	}

	// 循环遍历为了找到 链表的尾部。
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	// 将新的结点 插入到尾部。
	temp.next = newNode
	newNode.pre = temp
}

// 插入方式，按照节点中 no的升序顺序进行插入。
func InsertHeroNode2(head *HeroNode, newNode *HeroNode) {
	temp := head

	// 判断 头结点是否为空
	if temp == nil {
		fmt.Println("链表的表头为空， 无法进行链表插入操作。")
		return
	}

	// 遍历链表，找到合适的位置。
	flag := true
	for {
		if temp.next == nil {
			break
		} else if temp.next.no > newNode.no {
			break
		} else if temp.next.no == newNode.no {
			flag = false
		}

		temp = temp.next
	}

	if flag {
		newNode.pre = temp
		newNode.next = temp.next
		temp.next = newNode
		if temp.next != nil {
			temp.next.pre = newNode
		}
	} else {
		fmt.Println("当前结点 不符合插入条件。")
	}
}

// 根据结点的no，删除一个结点
func DelHeroNode(head *HeroNode, no int) {
	temp := head
	if temp == nil {
		fmt.Println("该 链表的表头为空，无法进行删除操作。")
		return
	}

	flag := false
	for {
		if temp.next == nil {
			break
		}

		if temp.next.no == no {
			flag = true
			break
		}
		temp = temp.next
	}

	if flag {
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.pre = temp
		}
		fmt.Println("要删除的结点，No= ", no, "删除成功。")
	} else {
		fmt.Println("要删除的结点，No= ", no, "不存在。")
	}
}

// 遍历显示 链表， 这个方法是采用正向的方式 遍历并显示链表的内容的。
func ListHeroNode(head *HeroNode) {
	temp := head
	if temp == nil {
		fmt.Println("该 链表的头为空，无法进行查看链表的操作。")
		return
	}

	for {
		if temp.next == nil {
			break
		}

		fmt.Printf("no:%d, name:%s, nickname:%s \n", temp.next.no, temp.next.name, temp.next.nickname)
		temp = temp.next
	}
}

// 遍历链表，并从尾部 逆向显示链表的内容。
func listHeroNodeByTail(head *HeroNode) {
	temp := head
	if temp == nil {
		fmt.Println("链表头为空， 无法进行 遍历显示操作。")
		return
	}

	// 将temp指针 移动到链表的尾部。
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	// 逆向遍历显示 链表的内容
	for {
		if temp.pre == nil {
			break
		}

		fmt.Printf("no:%d, name:%s, nickname:%s \n", temp.no, temp.name, temp.nickname)
		temp = temp.pre
	}
}

func main() {
	// 创建一个链表头。
	head := &HeroNode{}

	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}

	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}

	hero3 := &HeroNode{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
	}

	InsertHeroNode(head, hero1)
	InsertHeroNode(head, hero2)
	InsertHeroNode(head, hero3)
	ListHeroNode(head)
	fmt.Println("---------------------------------------")
	listHeroNodeByTail(head)
	fmt.Println("---------------------------------------")

	// 尝试加入已经存在的no的 结点， 结果是失败。
	InsertHeroNode2(head, hero2)

	// 删除 一个结点，但是这个结点的no是不存在的，
	fmt.Println("-------------------删除结点 操作。--------------------")
	DelHeroNode(head, 1)
	fmt.Println("-------------------删除结点后的链表。--------------------")
	ListHeroNode(head)
}
