package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

var FEED_URL string = "http://b.hatena.ne.jp/entrylist.rss"

type HatenaBookmark struct {
	Title []string `xml:"item>title"`
	Link  []string `xml:"item>link"`
}

func main() {
	hb, err := getHatenaBookmark(FEED_URL)

	if err != nil {
		log.Fatalf("Log: %v", err)
		return
	}

	fmt.Println(hb.Title)
	for n, v := range hb.Link {
		if n > 0 {
			fmt.Printf("%s \n", v)
		}
	}
}

// getHatenaBookmark ははてブを読み取って構造体に入れる
func getHatenaBookmark(feed string) (p *HatenaBookmark, err error) {

	// feedのURLからXMLを取得
	res, err := http.Get(feed)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	// 構造体作成
	wh := new(HatenaBookmark)
	// XMLをパース
	err = xml.Unmarshal(b, &wh)
	log.Printf("%#v", wh)

	return wh, err
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
