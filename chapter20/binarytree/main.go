package main

type Hero struct {
	Name  string
	ID    int
	Right *Hero
	Left  *Hero
}

func main() {
	root := &Hero{
		Name: "宋江",
		ID:   1,
	}

	left := &Hero{
		Name: "吴勇",
		ID:   2,
	}
	right1 := &Hero{
		Name: "卢俊义",
		ID:   3,
	}

	root.Right = right1
	root.Left = left

	right2 := &Hero{
		Name: "林冲",
		ID:   4,
	}
}
