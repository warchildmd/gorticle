package parser

import (
	"golang.org/x/net/html"
	"math"
)

var (
	ignoredTags = [...]string{"script", "style", "noscript"}
	ignoredTagPenalty = -1024.0
	depthPenalty = 0.5
)

func findBody(root *html.Node) *html.Node {
	for c := root.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == "body" {
			return c
		}
		x := findBody(c)
		if x != nil {
			return x
		}
	}
	return nil
}

func getBestContentNode(root *html.Node) *html.Node {
	var maxScore = 0.0
	var bestNode *html.Node

	var traverse func(*html.Node) float64
	traverse = func(n *html.Node) float64 {
		if isIgnoredTag(n.Data) {
			return ignoredTagPenalty
		}

		score := 0.0
		if n.Type == html.TextNode {
			score += calculateScore(n.Data)
		}

		childrenScore := 0.0
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			childScore := traverse(c)
			if c.Type != html.TextNode {
				childScore *= depthPenalty
			}
			childrenScore += childScore
		}
		score += childrenScore

		if score > maxScore {
			maxScore = score
			bestNode = n
		}

		return score
	}
	traverse(root)
	return bestNode
}

func calculateScore(text string) float64 {
	score := 0.0
	text = cleanText(text)
	if len(text) > 0 {
		score += math.Pow(float64(len(text)), 1.2)
	}
	return score
}



func getTextContentFromNode(n *html.Node) string {
	content := ""

	if isIgnoredTag(n.Data) {
		return content
	}

	if n.Type == html.TextNode {
		cleanContent := cleanText(n.Data) // strings.TrimLeft(n.Data, "\n\t ") // strings.TrimLeft(strings.Trim(n.Data, "\n\t"), " ")
		if len(cleanContent) > 0 {
			content += cleanContent
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		content += getTextContentFromNode(c)
	}

	return content
}

func isIgnoredTag(tag string) bool {
	for _, element := range ignoredTags {
		if element == tag {
			return true
		}
	}
	return false
}