package liuli

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

// Articles container of Article
type Articles struct {
	Articles []Article `json:"articles"`
}

// Article single liuli article
type Article struct {
	Description string `json:"description"`
	Img         string `json:"img"`
	Link        string `json:"link"`
	Title       string `json:"title"`
}

// GetArticles get articles from hacg.me
func GetArticles(page int) *Articles {
	uri := "https://www.hacg.me/wp/"
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
	articles.Articles = GetArticleArray(doc)
	return articles
}

// GetArticleArray get Article Objects from goquery document
func GetArticleArray(doc *goquery.Document) []Article {
	var articles []Article
	doc.Find("article").Each(func(index int, selection *goquery.Selection) {
		var tmp Article
		tmp.Description = selection.Find(".entry-content").Text()
		tmp.Img, _ = selection.Find("img").Attr("src")
		tmp.Link, _ = selection.Find(".more-link").Attr("href")
		tmp.Title = selection.Find(".entry-title").Text()
		if tmp.Title != "" {
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
