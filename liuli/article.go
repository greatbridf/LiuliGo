package liuli

import (
	"github.com/pkg/errors"
	"sync"

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
	var (
		articles  Articles
		mux       sync.Mutex
		waitGroup sync.WaitGroup
		cache     Cache
	)
	err := cache.Init()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer cache.Close()

	selections := doc.Find("article")
	waitGroup.Add(len(selections.Nodes))
	chError := make(chan error, len(selections.Nodes))

	selections.Each(func(index int, selection *goquery.Selection) {
		go func(index int, selection *goquery.Selection) {
			defer waitGroup.Done()

			title := selection.Find(".entry-title").Text()
			if title == "" {
				return
			}

			description := selection.Find(".entry-content").Text()
			next_link, _ := selection.Find(".more-link").Attr("href")
			img_link, _ := selection.Find("img").Attr("src")

			if !cache.Find(img_link) {
				data, err := ReadFromURI(img_link)
				if err != nil {
					chError <- err
					return
				}

				err = cache.Add(img_link, data)
				if err != nil {
					chError <- err
					return
				}
			}

			article := Article{
				description,
				"https://static.greatbridf.top/liuli/" + cache.GetHash(img_link),
				next_link,
				title,
			}

			mux.Lock()
			articles = append(articles, article)
			mux.Unlock()
		}(index, selection)
	})

	waitGroup.Wait()
	close(chError)

	err, ok := <-chError
	if ok && err != nil {
		return nil, err
	}

	return articles, nil
}
