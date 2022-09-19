package main

import (
	"github.com/gocolly/colly"
	"github.com/nats-io/nats.go"
	"net/url"
	"os"
	"regexp"
	"time"
)

// nats 是 Go 实现的一个高性能分布式消息队列，适用于高并发高吞吐量的消息分发场景。
// 早期的 nats 以速度为重，没有支持持久化。从 16 年开始，nats 通过 nats-streaming 支持基于日志的持久化，以及可靠的消息传输。

// 结合 nats 和 colly 的消息生产

var  domain2Collector = map[string]*colly.Collector{}
var nc *nats.Conn
var maxDepth = 10
var natsURL = "nats://localhost:4222"

func factory(urlStr string) *colly.Collector {
	u, _ := url.Parse(urlStr)
	return domain2Collector[u.Host]
}

// 我们认为匹配该模式的是该网站的详情页
var detailRegex, _ = regexp.Compile(`/go/go\?p=\d+$`)
// 匹配下面模式的是该网站的列表页
var listRegex, _ = regexp.Compile(`/t/\d+#\w+`)

func initABCDECollector() *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains("www.abcdefg.com"),
		colly.MaxDepth(maxDepth),
	)

	c.OnResponse(func(resp *colly.Response) {
		// 做一些爬完之后的善后工作
		// 比如页面已爬完的确认存进 MySQL
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// 基本的反爬虫策略
		link := e.Attr("href")
		time.Sleep(time.Second * 2)

		// 正则 match 列表页的话，就 visit
		if listRegex.Match([]byte(link)) {
			c.Visit(e.Request.AbsoluteURL(link))
		}
		// 正则 match 落地页的话，就发消息队列
		if detailRegex.Match([]byte(link)) {
			nc.Publish("tasks", []byte(link))
			nc.Flush()
		}
	})
	return c
}

func initHIJKLCollector() *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains("www.hijklmn.com"),
		colly.MaxDepth(maxDepth),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	})

	return c
}

func init() {
	domain2Collector["www.abcdefg.com"] = initABCDECollector()
	domain2Collector["www.hijklmn.com"] = initHIJKLCollector()
	var err error
	nc, err = nats.Connect(natsURL)
	if err != nil {os.Exit(1)}
}

func main() {
	urls := []string{"https://www.abcdefg.com", "https://www.hijklmn.com"}
	for _, url := range urls {
		instance := factory(url)
		instance.Visit(url)
	}
}
