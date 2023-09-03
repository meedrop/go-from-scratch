package main

import (
	"fmt"
)

func main() {
	str := "Go is beautifule language"
	fmt.Printf("The length of str is: %d\n", len(str))

	// for循环结构
	for i := 0; i < len(str); i++ {
		fmt.Printf("Charter is: %c\n", str[i])
	}

	// continue 关键字只能用于for循环中
}
