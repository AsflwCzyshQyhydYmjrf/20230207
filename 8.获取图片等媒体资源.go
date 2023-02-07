package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://www.example.com/image.jpg"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}
