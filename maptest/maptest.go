package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// 创建一个动态map方法
	map1 := make(map[string]int)
	map1["qqq"] = 1
	map1["ccc"] = 3
	fmt.Println("map data:", map1)

	// 将map转成json字符串
	data := map[string]interface{}{
		"name":   "Alice",
		"Age":    18,
		"isMale": false,
	}
	jsonString, err := json.Marshal(data)
	if err != nil {
		fmt.Println("JSON编码出错:", err)
		return
	}

	// 打印JSON字符串
	fmt.Println("JOSN 字符串：", string(jsonString))

}
