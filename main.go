package main

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
)

var feedURL = "http://b.hatena.ne.jp/entrylist.rss"

// HatenaBookmark is.
type HatenaBookmark struct {
	Title []string `xml:"item>title"`
	Link  []string `xml:"item>link"`
}

func main() {

	hb, err := getHatenaBookmark(feedURL)

	if err != nil {
		log.Fatalf("Log: %v", err)
		return
	}

	HatebuInsert(hb)
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

	return wh, err
}

// HatebuInsert db insert.
func HatebuInsert(hb *HatenaBookmark) {
	// 第2引数の形式は "user:password@tcp(host:port)/dbname"
	db, err := sql.Open("mysql", "root:password@/gocrawler?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println(hb.Title)
	for i, v := range hb.Link {
		if i >= 0 {

			fmt.Printf("%d %s %s \n", i, v, hb.Title[i])
			// 引数付きでInsert文発行
			result, err := db.Exec(`
      INSERT INTO hatebu(title, link) VALUES(?, ?) ON DUPLICATE KEY UPDATE
	  link = ?;`, hb.Title[i], v, v)

			if err != nil {
				panic(err.Error())
			}
			// 影響を与えた件数を取得
			n, err := result.RowsAffected()
			if err != nil {
				panic(err.Error())
			}
			fmt.Println(n)

		}
	}


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
