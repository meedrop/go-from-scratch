package main

import (
	"fmt"
	"strings"

	"example.com/gotest/trans"
)

var twoPi = 2 * trans.Pi

var (
	brother string
	count   int
	choice  bool
)

func main() {
	// fmt.Printf("%s", runtime.Version())
	// var x = "ttt"
	// var y = "bbbb"
	// code, data := abc(x, y)
	abc()
	fmt.Println(twoPi)
}

func abc() {
	// var goos string = runtime.GOOS
	// fmt.Printf("The operation system is: %s|%s\n", goos, goos)

	// path := os.Getenv("PATH")
	// fmt.Printf("Path is %s\n", path)

	// var kkk string
	// kkk = path // 变量赋值，实际指向同一个内存
	// fmt.Printf("kkk Path is %s\n", kkk)

	// var cC float64 = 0.33333333
	// fmt.Println(cC)

	s := "hello world"
	fmt.Println(string(s[1]))

	news := strings.Replace(s, "hello", "baby", -1)
	fmt.Println(news)

}

/*
abc测试函数
*/
func Aabc(x float64, y int) (int, string) {
	// var goos string = runtime.GOOS
	// fmt.Printf("The operation system is: %s|%s\n", goos, goos)

	// path := os.Getenv("PATH")
	// fmt.Printf("Path is %s\n", path)

	// var kkk string
	// kkk = path // 变量赋值，实际指向同一个内存
	// fmt.Printf("kkk Path is %s\n", kkk)

	// var cC float64 = 0.33333333
	// fmt.Println(cC)
	fmt.Println("%s, %s", x, y)
	var msg string = "ok"
	return 0, msg
}
