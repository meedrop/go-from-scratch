package main

import (
	"fmt"
	"os"
)

func main() {
	// if-else结构
	bool1 := false
	if bool1 {
		fmt.Println("The value is true")
	} else if bool1 == false {
		fmt.Println("The value is false")
	} else {
		fmt.Println("no way")
	}

	// 与函数结合常用
	OsCheck()

}

// func OsCheck(filename string) interface{
// 	f, err := os.Open("a.txt")
// 	return f, err
//     //func Aabc(x float64, y int) (int, string) {
// }

func OsCheck() {
	f, err := os.Open("a.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer f.Close()
	//func Aabc(x float64, y int) (int, string) {
}
