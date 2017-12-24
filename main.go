package main

import (
	"log"
	"gorticle/parser"
)

func main() {
	var article *parser.Article

	article, _ = parser.ParseUrl("https://www.bloomberg.com/news/articles/2017-12-23/bitcoin-climbs-finding-floor-after-worst-selloff-since-2015")
	log.Print(article.Content)

	article, _ = parser.ParseUrl("http://www.telegraph.co.uk/news/2017/12/23/existence-ufos-proved-beyond-reasonable-doubt-says-former-pentagon/")
	log.Print(article.Content)

	article, _ = parser.ParseUrl("https://www.washingtonpost.com/news/politics/wp/2017/03/17/how-much-is-donald-trumps-travel-and-protection-costing-anyway/?utm_term=.bb14fde4a60b")
	log.Print(article.Content)
}
