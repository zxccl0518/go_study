package model

import ("fmt")

type Person struct {
	Name string
	age int
	sal float64
}

// 写一个工厂模式的函数，相当于构造函数。
func NewPerson(name string)  *Person{
	return &Person{
		Name:name,
	}
}

// 为了访问年龄 和 sal，我们变异而一堆setXxx的方法 和 getXxx的方法。
func (p *Person)SetAge(age int){
	if age >0 && age <150{
		p.age = age
	}else{
		fmt.Println("年龄范围不对.")
	}
}

func (p *Person)GetAge() int{
	return p.age
}

func (p *Person) GetSal() float64 {
	return p.sal
}

func (p *Person)SetSal(sal float64){
	if sal >= 3000 && sal <= 30000 {
		p.sal = sal
	}else{
		fmt.Println("薪水范围 有问题。")
	}
}
