package parser

import (
	"strings"
	"golang.org/x/net/html"
	"net/http"
	"io/ioutil"
	"log"
)

type Article struct {
	Title, Content, Description, Url, Image string
}

func ParseUrl(url string) (*Article, error) {
	body, err := readUrl(url)
	if err != nil {
		return nil, err
	}

	return Parse(body)
}

func Parse(body string) (*Article, error) {
	root, err := html.Parse(strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	content := getContent(root)
	metaMap := getMetaMap(root)

	return &Article{metaMap["title"], content, metaMap["description"],
		metaMap["url"], metaMap["image"]}, nil
}

func readUrl(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	// Close the connection
	defer response.Body.Close()

	bodyByteArray, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	body := string(bodyByteArray)

	return body, nil
}

func cleanText(text string) string {
	return strings.TrimLeft(text, "\n\t ")
}

func find(root *html.Node, tag string) *html.Node {
	for c := root.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == tag {
			return c
		}
		x := find(c, tag)
		if x != nil {
			return x
		}
	}
	return nil
}
