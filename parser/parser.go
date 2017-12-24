package parser

import (
	"strings"
	"golang.org/x/net/html"
	"net/http"
	"io/ioutil"
	"log"
)

type Article struct {
	Title, Content string
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

	bodyNode := findBody(root)
	if bodyNode != nil {
		root = bodyNode
	}
	content := getBestContentNode(root)

	contentText := getTextContentFromNode(content)

	return &Article{"", contentText}, nil
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