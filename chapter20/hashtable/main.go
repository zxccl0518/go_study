package main

import "fmt"

// 雇员的结构体
type Emp struct {
	ID   int
	Name string
	Next *Emp
}

// 雇员链表。
type EmpLink struct {
	Head *Emp
}

// 哈希散列,定义一个hashtable ，含有一个链表数组。
type HastTable struct {
	LinkArr [7]EmpLink
}

// 向散列中 插入一个雇员。
func (this *HastTable) Insert(emp *Emp) {
	index := this.HashFun(emp.ID)
	this.LinkArr[index].Insert(emp)
}

// 向 链表中插入一个雇员。[升序的方式插入]
func (this *EmpLink) Insert(emp *Emp) {
	if emp == nil {
		fmt.Println("雇员 信息错误。无法插入。")
		return
	}

	if this.Head == nil {
		this.Head = emp
		return
	}

	pre := this.Head

	// 如果当前准备插入的雇员 小于链表头的雇员id，则更换头
	if emp.ID < this.Head.ID {
		emp.Next = this.Head
		this.Head = emp
		return
	}

	for {
		if pre.Next == nil {
			// emp.Next = pre.Next
			// pre.Next = emp
			break
		}

		if emp.ID >= pre.Next.ID {
			// pre.Next = emp
			// emp.Next = pre.Next
			break
		}
		pre = pre.Next
	}

	emp.Next = pre.Next
	pre.Next = emp
}

// 打印哈希散列
func (this *HastTable) ListHash() {
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].ListLink(i)
	}
}

// 打印链表。
func (this *EmpLink) ListLink(arrIndex int) {
	head := this.Head
	for {
		if head == nil {
			break
		}

		fmt.Printf("arrLink[%d] = %d,%s\t", arrIndex, head.ID, head.Name)
		head = head.Next
	}
	fmt.Println()
}

// 编写一个hash散列方法
func (this *HastTable) HashFun(id int) int {
	// 取模7 是因为，hashtable 数组有7个元素，通过取模 算出应该往哪个数组中插入
	return id % 7
}

func (this *HastTable) FindByID(id int) (emp *Emp, res bool) {
	index := this.HashFun(id)
	emp, res = this.LinkArr[index].FindByID(id)

	return
}

func (this *EmpLink) FindByID(ID int) (emp *Emp, res bool) {
	head := this.Head
	res = false
	emp = nil

	for {
		if head == nil {
			break
		}
		if ID == head.ID {
			emp = head
			res = true
			break
		}
		head = head.Next
	}

	return
}

func main() {
	emp1 := &Emp{
		ID:   1,
		Name: "我是1",
	}
	emp2 := &Emp{
		ID:   2,
		Name: "我是1",
	}
	emp3 := &Emp{
		ID:   3,
		Name: "我是3",
	}
	emp4 := &Emp{
		ID:   4,
		Name: "我是4",
	}
	emp5 := &Emp{
		ID:   5,
		Name: "我是5",
	}
	emp6 := &Emp{
		ID:   6,
		Name: "我是6",
	}
	emp7 := &Emp{
		ID:   7,
		Name: "我是7",
	}
	emp8 := &Emp{
		ID:   8,
		Name: "我是8",
	}
	emp9 := &Emp{
		ID:   9,
		Name: "我是9",
	}
	emp15 := &Emp{
		ID:   15,
		Name: "我是15",
	}

	var hashtable HastTable
	hashtable.Insert(emp15)
	hashtable.Insert(emp8)
	hashtable.Insert(emp9)
	hashtable.Insert(emp1)
	hashtable.Insert(emp2)
	hashtable.Insert(emp3)
	hashtable.Insert(emp4)
	hashtable.Insert(emp5)
	hashtable.Insert(emp6)
	hashtable.Insert(emp7)

	hashtable.ListHash()

	emper, res := hashtable.FindByID(15)
	if res == true {
		fmt.Printf("emper name = %s, id = %d\n", emper.Name, emper.ID)
	} else {
		fmt.Println("对应的id 不存在。")
	}
}
