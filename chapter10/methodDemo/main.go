package main

import "fmt"

type Person struct {
	Name string
}

// 给A类型绑定一份方法。
func (a Person) test() {
	fmt.Println("test()", a.Name)
}

func main() {
	person := Person{"刘亦菲"}
	person.test()
}
