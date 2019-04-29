package main

import "fmt"

// 二叉树结构体。
type Hero struct {
	Name        string
	ID          int
	Left, Right *Hero
}

// preOrder 前序遍历。所謂前序遍历就是边遍历跟结点，然后遍历其左结点，然后再遍历其右结点。
func preOrder(node *Hero) {
	if node != nil {
		fmt.Printf("node name = %v, id = %d \n", node.Name, node.ID)
		preOrder(node.Left)
		preOrder(node.Right)
	}
}

// 中序遍历的方式 遍历二叉树。
func Middle(node *Hero) {
	if node != nil {
		Middle(node.Left)
		fmt.Printf("中序遍历的方式。 node name = %v, id = %d \n", node.Name, node.ID)
		Middle(node.Right)
	}
}

// 中序遍历的方式 遍历二叉树。
func PostOrder(node *Hero) {
	if node != nil {
		PostOrder(node.Left)
		PostOrder(node.Right)
		fmt.Printf("后续 遍历的方式。 node name = %v, id = %d \n", node.Name, node.ID)
	}
}

func main() {
	root := &Hero{
		Name: "宋江",
		ID:   1,
	}

	left1 := &Hero{
		Name: "吴用",
		ID:   2,
	}
	root.Left = left1

	right1 := &Hero{
		Name: "卢俊义",
		ID:   3,
	}
	root.Right = right1

	right2 := &Hero{
		Name: "林冲",
		ID:   4,
	}
	right1.Right = right2

	tom := &Hero{
		Name: "tom",
		ID:   10,
	}
	left1.Left = tom

	jack := &Hero{
		Name: "jack",
		ID:   12,
	}
	left1.Right = jack

	// preOrder(root)
	// Middle(root)
	PostOrder(root)
}
