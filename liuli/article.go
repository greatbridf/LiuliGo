package liuli

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

// Articles container of Article
type Articles struct {
	articles []Article
}

// Article single liuli article
type Article struct {
	description string
	img         string
	link        string
	title       string
}

// GetArticles get articles from hacg.me
func GetArticles(page int) *Articles {
	uri := "https://www.liuli.in/wp/"
	if page != 1 {
		uri = "https://www.hacg.me/wp/page/" + strconv.Itoa(page)
	}
	res, err := http.Get(uri)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		panic(res.StatusCode)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}
	articles := new(Articles)
	articles.articles = GetArticleArray(doc)
	return articles
}

// GetArticleArray get Article Objects from goquery document
func GetArticleArray(doc *goquery.Document) []Article {
	var articles []Article
	doc.Find("article").Each(func(index int, selection *goquery.Selection) {
		var tmp Article
		tmp.description = selection.Find(".entry-content").Text()
		tmp.img, _ = selection.Find("img").Attr("src")
		tmp.link, _ = selection.Find(".more-link").Attr("href")
		tmp.title = selection.Find(".entry-title").Text()
		if tmp.title != "" {
			articles = append(articles, tmp)
		}
	})
	return articles
}

// GetArticlesJSON convert Articles Objecet into json string
func GetArticlesJSON(articles *Articles) string {
	jsonByteArray, err := json.Marshal(articles)
	if err != nil {
		panic(err)
	}
	return string(jsonByteArray)
}
