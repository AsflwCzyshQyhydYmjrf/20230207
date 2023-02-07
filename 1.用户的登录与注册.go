package main

import (
	"fmt"
	"strings"
)

// 用户信息
type User struct {
	Name string
	Pwd  string
}

// 用户信息map
var userMap map[string]User

// 初始化
func init() {
	userMap = make(map[string]User)
}

// 用户注册
func Register(name string, pwd string) {
	userMap[name] = User{name, pwd}
	fmt.Println("注册成功！")
}

// 用户登录
func Login(name string, pwd string) {
	user, ok := userMap[name]
	if ok {
		if strings.EqualFold(user.Pwd, pwd) {
			fmt.Printf("欢迎回来，%s\n", user.Name)
		} else {
			fmt.Println("登录失败，密码错误！")
		}
	} else {
		fmt.Println("登录失败，用户不存在！")
	}
}

func main() {
	var name string
	var pwd string
	fmt.Print("请输入用户名：")
	fmt.Scanln(&name)
	fmt.Print("请输入密码：")
	fmt.Scanln(&pwd)
	Login(name, pwd)
	Register(name, pwd)
}
