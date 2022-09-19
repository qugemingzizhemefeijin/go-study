package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
	"time"
)

// 《Go 语言编程》一书给出了简单的爬虫示例，经过了多年的发展，现在使用 Go 语言写一个网站的爬虫要更加方便，
// 比如用 colly 来实现爬取某网站（虚拟站点，这里用 abcdefg 作为占位符）在 Go 语言标签下的前十页内容。

var visited = map[string]bool{}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.abcdefg.com"),
		colly.MaxDepth(1),
	)

	// 我们认为匹配该模式的是该网站的详情页
	detailRegex, _ := regexp.Compile(`/go/go\?p=\d+$`)
	// 匹配下面模式的是该网站的列表页
	listRegex, _ := regexp.Compile(`/t/\d+#\w+`)

	// 所有 a 标签，上设置回调函数
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		// 已访问过的详情页或列表页，跳过
		if visited[link] && (detailRegex.Match([]byte(link)) || listRegex.Match([]byte(link))) {
			return
		}

		// 既不是列表页，也不是详情页
		// 那么不是我们关心的内容，要跳过
		if !detailRegex.Match([]byte(link)) && !listRegex.Match([]byte(link)) {
			println("not match", link)
			return
		}

		// 因为大多数网站有反爬虫策略
		// 所以爬虫逻辑中应该有 sleep 逻辑以避免被封杀
		time.Sleep(time.Second)
		println("match", link)

		visited[link] = true

		time.Sleep(time.Millisecond * 2)
		c.Visit(e.Request.AbsoluteURL(link))
	})

	err := c.Visit("https://www.abcdefg.com/go/go")
	if err != nil {fmt.Println(err)}
}
