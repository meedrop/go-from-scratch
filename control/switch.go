package main

import (
	"fmt"
)

func main() {
	num1 := 7
	switch {
	case num1 < 0:
		fmt.Println("Number is negative")
	case num1 > 0 && num1 < 10:
		fmt.Println("Number between 0 to 10")
	default:
		fmt.Println("Number is 10 or greater")
	}
}
