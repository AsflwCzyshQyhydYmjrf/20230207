package main

import (
	"fmt"
)

type BookReview struct {
	title  string
	author string
	review string
}

func main() {
	// 创建一个书评
	bookReview := BookReview{title: "The Great Gatsby", author: "F. Scott Fitzgerald", review: "A classic of American literature"}

	// 打印书评
	fmt.Println(bookReview)

	// 修改书评
	bookReview.review = "A timeless classic"

	// 打印修改后的书评
	fmt.Println(bookReview)

	// 删除书评
	bookReview = BookReview{}

	// 打印删除后的书评
	fmt.Println(bookReview)
}
