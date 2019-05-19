package liuli

import (
	"bytes"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"

	"crypto/md5"
	"encoding/hex"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

// RenderHTMLTag convert goquery.Selection to html string
func RenderHTMLTag(selection *goquery.Selection) string {
	node := selection.Get(0)
	buf := bytes.NewBuffer([]byte{})
	html.Render(buf, node)
	return buf.String()
}

func ReadFromURI(uri string) ([]byte, error) {
	res, err := http.Get(uri)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get response")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read the stream")
	}
	return body, nil
}

func Hash(input []byte) string {
	m := md5.New()
	m.Write(input)
	output := hex.EncodeToString(m.Sum(nil))
	return output
}
