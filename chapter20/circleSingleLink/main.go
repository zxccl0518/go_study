package main

import "fmt"

type CatNode struct {
	no   int
	name string
	next *CatNode
}

// 向 环形链表中插入元素
func InsertCatNodeLink(head *CatNode, cat *CatNode) {
	if head == nil {
		return
	}

	// 如果next为空，正明将要插入的结点 将会是头结点。
	if head.next == nil {
		head.no = cat.no
		head.name = cat.name
		head.next = head
		fmt.Println("结点 ", cat, "加入 环形单向链表 成功。")
		return
	}

	temp := head
	for {
		// temp是环形单向链表的最后一个结点。
		if temp.next == head {
			break
		}
		temp = temp.next
	}

	// 将花朵插入到环形链表的最后一个位置上。
	temp.next = cat
	cat.next = head
}

func InsertNod(head *CatNode, newNode *CatNode) {
	if head.next == nil {
		head.no = newNode.no
		head.name = newNode.name
		head.next = head
		return
	}

	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	temp.next = newNode
	newNode.next = head
}

func ListCatLink(head *CatNode) {
	if head == nil {
		fmt.Println("ListCatLink() --- 环形链表为空。")
		return
	}

	temp := head
	for {
		fmt.Printf("cat no = %d, name = %s \n", temp.no, temp.name)

		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

// 按照升序排列。
func InsertCatNodeLinkByNo(head *CatNode, cat *CatNode) (newHead *CatNode) {
	// 链表 头为空指针
	if head == nil {
		fmt.Println("InsertCatNodeLinkByNo() --- 环形链表 头指针为空")
		return
	}

	temp := head
	newHead = head
	if head.next == nil {
		head.no = cat.no
		head.name = cat.name
		head.next = head
		return
	}

	// 链表的头 发生改变。
	if temp.no > cat.no {
		// 先找到最后一个结点指针。
		for {
			if temp.next == head {
				break
			}
			temp = temp.next
		}
		cat.next = head
		head = cat
		temp.next = head
		newHead = cat
		return
	}

	for {
		// 如果到了最后一个,说明前面没有比这个新加入的结点还大的。
		if temp.next == head {
			temp.next = cat
			cat.next = head
			break
		}

		// 从小到大 插入结点。
		if temp.next.no > cat.no {
			cat.next = temp.next
			temp.next = cat
			break
		}

		temp = temp.next
	}

	return
}

// 可以 删除任意结点。
func DelCatNode(head *CatNode, id int) (newHead *CatNode) {
	if head.next == nil {
		fmt.Println("链表的头指针为空，删除操作失败")
		return
	}

	temp := head
	newHead = head

	// 如果头结点 就是要删除的那个节点的情况。
	if head.no == id {
		fmt.Println("找到了 要删除的结点 id = ", temp.no)
		// 该链表就一个结点。
		if head.next == head {
			newHead = &CatNode{}
		} else {
			for {
				if temp.next == head {
					break
				}
				temp = temp.next
			}
			newHead = newHead.next
			temp.next = newHead
		}
	} else {
		// 要删除的结点不在头节点的位置。
		for {
			if temp.next.no == id {
				fmt.Println("找到了 要删除的结点 id = ", temp.next.no)
				break
			}
			temp = temp.next
		}
		temp.next = temp.next.next
	}

	return
}

func main() {
	catHead := &CatNode{}

	cat1 := &CatNode{
		no:   1,
		name: "cat1",
	}
	cat2 := &CatNode{
		no:   2,
		name: "cat2",
	}
	cat3 := &CatNode{
		no:   5,
		name: "cat3",
	}

	cat4 := &CatNode{
		no:   7,
		name: "cat7",
	}
	cat5 := &CatNode{
		no:   3,
		name: "cat3",
	}

	cat6 := &CatNode{
		no:   0,
		name: "cat6",
	}

	cat7 := &CatNode{
		no:   100,
		name: "cat7",
	}

	cat8 := &CatNode{
		no:   10,
		name: "cat8",
	}
	// InsertCatNodeLink(catHead, cat2)
	// InsertCatNodeLink(catHead, cat3)
	// InsertCatNodeLink(catHead, cat1)

	catHead = InsertCatNodeLinkByNo(catHead, cat2)
	catHead = InsertCatNodeLinkByNo(catHead, cat3)
	catHead = InsertCatNodeLinkByNo(catHead, cat1)
	catHead = InsertCatNodeLinkByNo(catHead, cat4)
	catHead = InsertCatNodeLinkByNo(catHead, cat5)
	catHead = InsertCatNodeLinkByNo(catHead, cat6)
	catHead = InsertCatNodeLinkByNo(catHead, cat7)
	catHead = InsertCatNodeLinkByNo(catHead, cat8)
	ListCatLink(catHead)

	catHead = DelCatNode(catHead, 0)
	catHead = DelCatNode(catHead, 1)
	catHead = DelCatNode(catHead, 2)
	catHead = DelCatNode(catHead, 100)
	ListCatLink(catHead)
}
