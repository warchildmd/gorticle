package parser

import "golang.org/x/net/html"

func getMetaMap(root *html.Node) map[string]string {
	result := make(map[string]string)

	headNode := find(root, "head")
	if headNode == nil {
		return result
	}

	for c := headNode.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == "meta" {
			attributes := c.Attr
			content := ""
			field := ""
			for _, element := range attributes {
				if element.Key == "property" {
					switch element.Val {
					case "og:url":
						field = "url"
						break
					case "og:title":
						field = "title"
						break
					case "og:description":
						field = "description"
						break
					case "og:image":
						field = "image"
						break
					}
				}
				if element.Key == "content" {
					content = element.Val
				}
			}
			result[field] = content
		}
	}

	_, titleExtracted := result["title"]
	if titleExtracted == false {
		result["title"] = "" // TODO Extract title from the <title> tag
	}

	return result
}
