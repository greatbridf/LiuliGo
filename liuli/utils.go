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
		Log.E(err.Error())
		return nil, errors.Wrap(err, "Cannot get response")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		Log.E(err.Error())
		return nil, errors.Wrap(err, "Cannot read from stream reader")
	}
	return body, nil
}

func HE(err error) error {
	Log.E(err.Error())
	return err
}

func HEM(err error, msg string) error {
	errors.Wrap(err, msg)
	Log.E(err.Error())
	return err
}

func Hash(input []byte) string {
	m := md5.New()
	m.Write(input)
	output := hex.EncodeToString(m.Sum(nil))
	return output
}
