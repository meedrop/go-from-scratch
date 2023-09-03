package main

import (
	"fmt"
)

func main() {

	// *固定数组
	// 声明arr这个切片的长度是5，里面的数据是int
	var arr [5]int
	arr[0] = 33
	fmt.Println("arr data:", arr)

	// for i := 0; i < len(arr); i++ {
	// 	fmt.Printf("index: %d, value: %d\n", i, arr[i])
	// }

	// for-range方式生成
	for i, _ := range arr {
		fmt.Printf("index: %d, value: %d\n", i, arr[i])
	}

	// 可变数组
	var Carr []int

	Carr = append(Carr, 1)
	Carr = append(Carr, 2)
	Carr = append(Carr, 3)
	Carr = append(Carr, 4)
	Carr = append(Carr, 5)

	for i, _ := range Carr {
		fmt.Printf("index: %d, value: %d\n", i, Carr[i])
	}
}
