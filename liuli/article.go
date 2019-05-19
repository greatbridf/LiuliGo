package liuli

import (
	"github.com/pkg/errors"

	"github.com/PuerkitoBio/goquery"
)

// Article single liuli article
type Article struct {
	Description string `json:"description"`
	Img         string `json:"img"`
	Link        string `json:"link"`
	Title       string `json:"title"`
}

type Articles []Article

// GetArticles get articles from hacg.me
func GetArticles(page string) (Articles, error) {
	uri := "https://www.hacg.me/wp/"
	if page != "1" {
		uri = "https://www.hacg.me/wp/page/" + page
	}
	doc, err := goquery.NewDocument(uri)
	if err != nil {
		return nil, errors.Wrap(err, ERR_CANNOT_GOQUERY)
	}
	articles, err := GetArticleArray(doc)
	if err != nil {
		return nil, errors.Wrap(err, "Cannot get article array")
	}
	return articles, nil
}

// GetArticleArray get Article Objects from goquery document
func GetArticleArray(doc *goquery.Document) (Articles, error) {
	var articles Articles
	var ERR error
	cache := Cache{}
	err := cache.Init()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer cache.Close()
	doc.Find("article").Each(func(index int, selection *goquery.Selection) {
		if ERR != nil {
			return
		}
		tmp := Article{
			Description: selection.Find(".entry-content").Text(),
			Title:       selection.Find(".entry-title").Text(),
		}
		if tmp.Title == "" {
			return
		}
		tmp.Link, _ = selection.Find(".more-link").Attr("href")
		img_link, _ := selection.Find("img").Attr("src")
		if !cache.Find(img_link) {
			data, err := ReadFromURI(img_link)
			if err != nil {
				ERR = errors.Wrap(err, "Cannot get title image")
				return
			}
			err = cache.Add(img_link, data)
			if err != nil {
				ERR = errors.Wrap(err, "Cannot add title image to cache")
				return
			}
		}
		tmp.Img = "https://static.greatbridf.top/liuli/" + cache.GetHash(img_link)
		articles = append(articles, tmp)
	})
	if ERR != nil {
		return nil, errors.WithStack(ERR)
	}
	return articles, nil
}
