package binaryTreeExer

import "fmt"

type Item struct {
	key int
}

type Tree struct {
	lchild, rchild *trTreeee
	item           Item
	count          int
}

func Create(T *Tree, x Item) *Tree {
	if T == nil {
		T = new(Tree)
		T.item = x
		T.count = 1
	} else if compare(T.Item, x) == 1 {
		T.lchild = create(T.lchild, x)
	} else if compare(T.Item, x) == 0 {
		T.count++
	} else {
		T.rchild = create(T.rchild, x)
	}

	return T
}

// 比较二叉树的大小。
func Compare(x1 Item, x2 Item) int {
	ret := 0
	if x1.key > x2.key {
		ret = 1
	} else if x1.key == x2.key {
		ret = 0
	} else {
		ret = -1
	}
	return ret
}

// 使用前序遍历的方式。遍历二叉树。
func PreOrder(T *Tree) {
	if T != nil {
		fmt.Printf("T.item.key = %d \n", T.item.key)
		PreOrder(T.lchild)
		PreOrder(T.rchild)
	}
}
