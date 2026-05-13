package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector(colly.MaxDepth(2), colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36"))

	// c.Limit()

	// 每个请求之前执行的钩子函数
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	// 每个请求发送之后接收到响应时执行的钩子函数
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visiting response:", r.StatusCode)
	})

	// 当找到 a 标签时调用的回调函数
	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		link := h.Attr("href")
		if link != "" {
			fmt.Println("Found link:", link)
			h.Request.Visit(link)
		}
	})
	// 注销指定规则的注册的回调函数
	// c.OnHTMLDetach("a[href]")

	// 发生错误时的钩子函数
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	// 开始爬虫, 从指定 URL 开始
	err := c.Visit("https://www.autohome.com")
	if err != nil {
		log.Fatalln(err)
	}
	// 等待所有异步完成
	c.Wait()
}
