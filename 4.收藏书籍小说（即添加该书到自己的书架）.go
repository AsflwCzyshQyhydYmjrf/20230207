//定义书籍结构体
type Book struct {
Name string
Author string
Pages int
}

//定义用户结构体
type User struct {
Name string
BookShelf []Book
}

//定义收藏书籍函数
func (u *User) CollectBook(b Book) {
u.BookShelf = append(u.BookShelf, b)
}

func main() {
//初始化用户
u := &User{
Name: "小明",
}
//初始化书籍
b := Book{
Name: "三国演义",
Author: "罗贯中",
Pages: 500,
}
//收藏书籍
u.CollectBook(b)
fmt.Println(u)
}