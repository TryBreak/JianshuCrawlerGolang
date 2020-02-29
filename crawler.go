package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*

 id 为  list-container

 下的

 ul标签  class="note-list"

 下的

 li标签

 下的

 class="title"  a 标签的内容

*/

func main() {
	fmt.Println("Hello, world")
	url := "http://www.jianshu.com"
	download(url)
}

func download(url string) {
	fmt.Println("Hello, world")
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http get error", err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error", err)
		return
	}

	writeString := string(body)

	// fmt.Println(writeString)

	writeTxt(writeString)
}

func writeTxt(s string) {
	content := []byte(s)
	ioutil.WriteFile("test.txt", content, 0644)
}
