package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
)

var url = "https://matome.naver.jp/odai/2138718423424640701" //テスト用url

func main() {

	flag.Parse()
	filename := flag.Arg(0)

	if err := os.Mkdir(filename, 0777); err != nil {
		fmt.Println("faild")
		os.Exit(1)
	}

	doc, _ := goquery.NewDocument(url)
	var i int
	doc.Find("div > p > a > img").Each(func(_ int, s *goquery.Selection) {
		text, _ := s.Attr("src")
		fmt.Println(text)
		err := SavePic(filename, text, i)
		if err != nil {
			fmt.Println("faild2")
			os.Exit(1)
		}
		i++
	})
}
