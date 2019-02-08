package liuli

import (
	"github.com/pkg/errors"

	"github.com/PuerkitoBio/goquery"
)

// GetContent get page by link
// Both content and styles
func GetContent(id string) (string, error) {
	cache := Cache{}
	err := cache.Init()
	if err != nil {
		return "", errors.WithStack(err)
	}
	defer cache.Close()
	if cache.Find(id) {
		data, err := cache.Get(id)
		if err != nil {
			return "", errors.WithStack(err)
		}
		return string(data), nil
	} else {
		link := "https://www.hacg.me/wp/" + id + ".html"
		doc, err := goquery.NewDocument(link)
		if err != nil {
			Log.E(err.Error())
			return "", errors.Wrap(err, ERR_CANNOT_GOQUERY)
		}

		content, err := GetContentNoStyle(doc)
		if err != nil {
			Log.E(err.Error())
			return "", errors.WithStack(err)
		}

		style := GetStyle(doc)
		data := style + content

		err = cache.Add(id, []byte(data))
		if err != nil {
			Log.E(err.Error())
			return "", errors.WithStack(err)
		}
		return data, nil
	}

}

// GetContentNoStyle get content from doc
func GetContentNoStyle(doc *goquery.Document) (string, error) {
	content := ""
	var ERR error
	doc.Find(".entry-content").Each(func(_ int, selection *goquery.Selection) {
		if ERR != nil {
			return
		}
		tmp, err := selection.Html()
		if err != nil {
			ERR = errors.WithStack(err)
			return
		}
		content += tmp
	})
	if ERR != nil {
		Log.E(ERR.Error())
		return "", ERR
	}
	return content, nil
}

// GetStyle get css tags from doc
func GetStyle(doc *goquery.Document) string {
	style := ""
	doc.Find("link[rel='stylesheet']").Each(func(_ int, selection *goquery.Selection) {
		tmp := RenderHTMLTag(selection)
		style += (tmp + "\n")
	})
	return style
}
