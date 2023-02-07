package main

import (
	"fmt"
)

func main() {
	// 定义书籍标签
	var bookTags = []string{"玄幻", "都市", "历史", "武侠", "科幻"}

	// 获取书籍标签
	var tag string
	fmt.Printf("请输入书籍标签：")
	fmt.Scanln(&tag)

	// 判断书籍标签是否存在
	for _, v := range bookTags {
		if v == tag {
			fmt.Println("获取书籍成功！")
			break
		}
	}
}
