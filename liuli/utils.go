package liuli

import (
	"bytes"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
    "crypto/md5"
    "encoding/hex"
)

// RenderHTMLTag convert goquery.Selection to html string
func RenderHTMLTag(selection *goquery.Selection) string {
	node := selection.Get(0)
	buf := bytes.NewBuffer([]byte{})
	html.Render(buf, node)
	return buf.String()
}

func Hash(input string) string {
    m := md5.New()
    m.Write([]byte(input))
    output := hex.EncodeToString(m.Sum(nil))
    return output
}

