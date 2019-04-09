package main

import "fmt"

// 定义一个heroNode struct
type HeroNode struct {
	no       int
	name     string
	nickname string
	next     *HeroNode // 表示指向下一个节点。
}

// 给链表插入一个节点。
// 编写一种插入方法，在单链表的最后加入。
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	// 思路
	// 1.先找到该链表的最后逇节点。
	// 2.创建一个辅助节点【跑龙套，帮忙。】

	temp := head
	for {
		if temp.next == nil {
			break
		}

		temp = temp.next // 让temp 不断的指向下一个节点。
	}

	// 将newHero 加入到链表的最后。
	temp.next = newHeroNode
}

// 给链表插入一个节点。
// 编写一种插入方法，根据no的编号从小到大插入。[实用性较高]
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	// 思路
	// 1.先找到适当的结点
	// 2.创建一个辅助节点【跑龙套，帮忙。】

	temp := head
	flag := true
	// 让插入结点的no 和temp的下一个节点的no比较。
	for {
		if temp.next == nil { // 说明到链表的最后了。
			break
		} else if temp.next.no > newHeroNode.no {
			//说明newHeroNode就应该插入到temp后面。
			break
		} else if temp.next.no == newHeroNode.no {
			// 说明我们的链表中已经有这个no，就不允许插入。
			fmt.Println("相等的的情况不允许插入。")
			flag = false
			break
		}

		temp = temp.next
	}
	if !flag {
		fmt.Println("对不起，已经存在，no - ", newHeroNode.no)
	} else {
		newHeroNode.next = temp.next
		temp.next = newHeroNode
	}
}

// 删除一个人节点。
func DelHeroNode(head *HeroNode, id int) {
	temp := head
	flag := false
	// 找到要删除的节点的no
	for {
		if temp.next == nil { // 说明到链表的最后了。
			break
		} else if temp.next.no == id {
			// 说明我们的链表中已经有这个no，就删除
			flag = true
			break
		}

		temp = temp.next
	}

	// 找到了对应id的那个几点。
	if flag {
		temp.next = temp.next.next
	}
}

// 显示链表的所有节点信息。
func listHeroNode(heroNode *HeroNode) {
	if heroNode == nil {
		fmt.Println("该链表是一个空链表。")
		return
	}

	fmt.Println("----------------------------------")
	tempHead := heroNode
	for {
		if tempHead.next == nil {
			break
		}
		fmt.Println(tempHead.next.name)
		tempHead = tempHead.next
	}
	fmt.Println("----------------------------------")
}

func main() {
	// 1.先创建一个头结点.
	head := &HeroNode{}

	// 2.创建一个新的HeroNode
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
		name:     "林聪",
		nickname: "豹子头",
	}

	hero4 := &HeroNode{
		no:       3,
		name:     "吴用",
		nickname: "智多星",
	}

	// InsertHeroNode(head, hero1)
	// InsertHeroNode(head, hero2)
	// InsertHeroNode(head, hero3)

	InsertHeroNode2(head, hero3)
	InsertHeroNode2(head, hero2)
	InsertHeroNode2(head, hero1)
	InsertHeroNode2(head, hero4)
	//宋江,卢俊义,林聪

	listHeroNode(head)

	DelHeroNode(head, 3)

	listHeroNode(head)
}
