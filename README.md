## gorticle - Web articles content extractor

At the moment the parser extracts only the article content (no title, or other key fields)

Example:

```$xslt
article, err = parser.ParseUrl("https://www.bloomberg.com/news/articles/2017-12-23/bitcoin-climbs-finding-floor-after-worst-selloff-since-2015")
if err == nil {
    fmt.Print(err)
} else {
    log.Print(article.Content)
}
```