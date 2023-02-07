package main

import (
	"fmt"
)

type Book struct {
	name   string
	author string
	price  float64
}

func main() {
	// 初始化书籍
	book1 := Book{name: "《三国演义》", author: "罗贯中", price: 20.0}
	book2 := Book{name: "《西游记》", author: "吴承恩", price: 30.0}
	book3 := Book{name: "《水浒传》", author: "施耐庵", price: 25.0}
	// 将书籍放入到slice
	books := []Book{book1, book2, book3}
	// 遍历slice
	for _, book := range books {
		fmt.Println(book)
	}
	// 根据作者查询书籍
	searchByAuthor("罗贯中", books)
	// 根据价格查询书籍
	searchByPrice(30.0, books)
}

// 根据作者查询书籍
func searchByAuthor(author string, books []Book) {
	for _, book := range books {
		if book.author == author {
			fmt.Printf("查询到作者为%s的书籍：%s\n", author, book.name)
			return
		}
	}
	fmt.Printf("未查询到作者为%s的书籍\n", author)
}

// 根据价格查询书籍
func searchByPrice(price float64, books []Book) {
	for _, book := range books {
		if book.price == price {
			fmt.Printf("查询到价格为%f的书籍：%s\n", price, book.name)
			return
		}
	}
	fmt.Printf("未查询到价格为%f的书籍\n", price)
}
