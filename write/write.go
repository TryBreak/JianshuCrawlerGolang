package main

import (
	"io/ioutil"
)

func main() {
	content := []byte("我爱我的祖国")
	ioutil.WriteFile("test.txt", content, 0644)
}
