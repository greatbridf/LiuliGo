package liuli

import (
	"encoding/json"
	"github.com/pkg/errors"

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
func GetArticles(page string) (*Articles, error) {
	uri := "https://www.hacg.me/wp/"
	if page != "1" {
		uri = "https://www.hacg.me/wp/page/" + page
	}
	doc, err := goquery.NewDocument(uri)
	if err != nil {
		Log.E(err.Error())
		return nil, errors.Wrap(err, ERR_CANNOT_GOQUERY)
	}
	tmp, err := GetArticleArray(doc)
	if err != nil {
		Log.E(err.Error())
		return nil, errors.Wrap(err, "Cannot get article array")
	}
	articles := &Articles{
		Articles: tmp,
	}
	return articles, nil
}

// GetArticleArray get Article Objects from goquery document
func GetArticleArray(doc *goquery.Document) ([]Article, error) {
	var articles []Article
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
		tmp.Img = "https://interface.greatbridf.top/liuli?req=resource&hash=" + cache.GetHash(img_link)
		articles = append(articles, tmp)
	})
	if ERR != nil {
		Log.E(ERR.Error())
		return nil, ERR
	}
	return articles, nil
}

// GetArticlesJSON convert Articles Objecet into json string
func GetArticlesJSON(articles *Articles) (string, error) {
	jsonByteArray, err := json.Marshal(articles)
	if err != nil {
		Log.E(err.Error())
		return "", errors.Wrap(err, "Cannot stringify JSON")
	}
	return string(jsonByteArray), nil
}
