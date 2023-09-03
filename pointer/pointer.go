package main

import "fmt"

func main() {

	i1 := 5
	fmt.Printf("An integer: %d, its locations in memory: %p\n", i1, &i1)

	// var myPointer *int
	var myPointer *int
	myPointer = &i1
	fmt.Printf("memory location is %p, value is %d\n", myPointer, *myPointer)

	// 改变值，看两个变量是否同时更新
	i1 = 10
	fmt.Printf("myPointer value: %d, i1 value: %d", *myPointer, i1)

	// 常量字面量没有指针
	// const i = 5
	// ptr := &i //error: cannot take the address of i
	// ptr2 := &10 //error: cannot take the address of 10
}
