package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/PuerkitoBio/goquery"
)

func FetchLink(url string) ([]string, error) {
	var urls []string
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		urls = append(urls, link)
		//fmt.Println(link)
	})
	return urls, nil
}

func main() {
	//url := "http://daoju.qq.com/"
	url := os.Args[1]
	if len(url) == 0 {
		url = "http://desk.zol.com.cn/meinv/"
	}
	urls, err := FetchLink(url)
	if err != nil {
		log.Fatal(err)
	}
	for _, u := range urls {
		fmt.Println(u)
	}
}