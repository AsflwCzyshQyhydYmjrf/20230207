package main

import (
	"fmt"
)

// 定义书籍结构体
type Book struct {
	title  string
	author string
	price  float64
}

// 定义用户结构体
type User struct {
	name     string
	password string
}

// 定义书评结构体
type Review struct {
	user    User
	book    Book
	comment string
	rating  int
}

// 定义获取书评函数
func getReview(review Review) {
	fmt.Printf("用户：%s\n书名：%s\n作者：%s\n价格：%.2f\n评论：%s\n评分：%d\n\n", review.user.name, review.book.title, review.book.author, review.book.price, review.comment, review.rating)
}

func main() {
	// 创建一本书
	book := Book{
		title:  "Go语言编程",
		author: "李四",
		price:  99.99,
	}
	// 创建用户
	user := User{
		name:     "张三",
		password: "123456",
	}
	// 创建书评
	review := Review{
		user:    user,
		book:    book,
		comment: "这本书写的很好，学习Go语言很有帮助",
		rating:  5,
	}
	// 获取书评
	getReview(review)
}
