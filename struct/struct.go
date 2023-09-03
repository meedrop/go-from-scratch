package main

import (
	"fmt"
	"strings"
)

type struct1 struct {
	i1  int
	f1  float64
	str string
}

// Person类定义 和 方法
type Person struct {
	firstName string
	lastName  string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
	fmt.Println("upPerson firstName:", p.firstName)
	fmt.Println("upPerson lastName:", p.lastName)
}

func (p *Person) addAllName() string {
	return p.firstName + p.lastName
}

func main() {
	ms := new(struct1) // 初始化一个结构体（对象），等同于ms := &struct1{10, 15.5, "Chris"}

	// 赋值
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "Chris"

	fmt.Println("i1 is:", ms.i1) // 获取值
	fmt.Println("f1 is:", ms.f1)
	fmt.Println("str is:", ms.str)
	fmt.Println("ms is:", ms) // ms实际是一个内存地址
	/*
		output
		i1 is: 10
		f1 is: 10
		str is: 10
		ms is: &{10 15.5 Marry}

	*/

	// 对新增的对象操作
	fmt.Println(" ")
	fmt.Println(" 对新增的对象操作")
	myPerson := new(Person)
	myPerson.firstName = "Lin"
	myPerson.lastName = "QiuQiu"

	upPerson(myPerson) // 这里不传&是因为new定义了
	fmt.Println("firstName is:", myPerson.firstName)
	fmt.Println("lastName is:", myPerson.lastName)

	// 定义执行对象的方法
	fmt.Println(myPerson.addAllName())

}
