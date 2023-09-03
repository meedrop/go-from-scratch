package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("In main()")
	go longWait()
	go shortWait()
	go rangePrint1()
	go rangePrint2()

	fmt.Println("About to sleep in main()")
	time.Sleep(10 * 1e9)
	fmt.Println("At the end of main()")
}

func longWait() {
	fmt.Println("Beginning longWait()")
	time.Sleep(5 * 1e9) // 2秒
	fmt.Println("End longWait()")
}

func shortWait() {
	fmt.Println("Beginning shortWait()")
	time.Sleep(2 * 1e9) // 5秒
	fmt.Println("End shortWait()")
}

func rangePrint1() {
	for i := 0; i < 30; i++ {
		fmt.Println(i)
		time.Sleep(1 * 1e9)
	}
}

func rangePrint2() {
	for i := 0; i < 30; i++ {
		fmt.Println(i)
		time.Sleep(2 * 1e9)
	}
}
