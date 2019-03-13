package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/PuerkitoBio/goquery"
)
// 接收string作为url
func fetch(url string) ([]string, error) {
	var urls []string
	resp, err := http.Get(url) // 调用http库，http.get直接去发送get请求，把url的html抓下来。
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // 读取完做关闭操作
	if resp.StatusCode != http.StatusOK { // 如果resp.StatusCode不等于OK（200）
		return nil, errors.New(resp.Status) // 返回错误，如果等于200继续往下走
	}
	// 返回doc处理错误
	doc, err := goquery.NewDocumentFromResponse(resp) // 从response中创建一个Document
	if err != nil {
		return nil, err
	}
	// doc里面直接Find image Each func，这里面做了匿名函数，匿名函数i int, s *goquery.Selection是固定的，也就是说Each必须接收这个函数
	doc.Find("img").Each(func(i int, s *goquery.Selection) { // 这个函数会进行调用，对于里面所有的image标签，img标签它会做一个处理
		link, _ := s.Attr("src") // 处理就是返回所有的image标签中的src属性
		fmt.Println(link) // 把src属性中的link获取下来
	})
	return urls, nil
}

func main() {
	//url := "http://daoju.qq.com/"
	url := os.Args[1] // 先读取第一个参数
	urls, err := fetch(url) // 调用fetch函数，获取这个URL中对应的所有图片的URL。获取到之后返回URL List
	if err != nil { // 判断是否有错误
		log.Fatal(err)
	}
	for _, u := range urls {
		fmt.Println(u) // 最后直接把URL list 打印出来
	}
}