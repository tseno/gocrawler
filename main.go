package main

import (
	"fmt"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	Machikon()
}

// Machikon は街コンサイトの一覧を出力する。
func Machikon() {

	doc, err := goquery.NewDocument("https://machicon.jp/areas/tokyo")
	if err != nil {
		fmt.Println(err)
	}

	u := url.URL{}
	u.Scheme = doc.Url.Scheme
	u.Host = doc.Url.Host

	// ページtitleの取得
	title := doc.Find("title").Text()
	fmt.Println(title)

	// 掲載イベントURL一覧を取得
	doc.Find("article").Each(func(i int, s *goquery.Selection) {

		// aタグを検索
		aTag := s.Find("a")

		// #js-event_1736556 > header > a > h2 > span
		// aタグのhrefのvalueを取得(相対パスが取得される)

		text1 := s.Find("header > a > h2 > span").Text()
		h, _ := aTag.Last().Attr("href")
		u.Path = h
		fmt.Println(text1)
		fmt.Println(u.String())
	})
}
