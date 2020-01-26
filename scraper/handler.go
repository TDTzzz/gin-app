package main

import "C"
import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"log"
)

const (
	Limit     = 5
	UserAgent = "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"
	//URL="https://movie.douban.com/top250?start=0&filter="
	URL  = "https://wh.lianjia.com/ershoufang/"
	HOST = "https://wh.lianjia.com"
)

func main() {
	c := colly.NewCollector(
		colly.Async(true),
		colly.UserAgent(UserAgent),
	)

	c.Limit(&colly.LimitRule{DomainGlob: "*.lianjia.*", Parallelism: Limit})

	//sonC := c.Clone()

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	//step 1 爬各个区域
	c.OnHTML(".position", func(e *colly.HTMLElement) {
		posDom := e.DOM.Find("dl").Eq(1).Find("div").Eq(0).Find("div").Find("a")
		log.Println(c.ID)
		posDom.Each(func(i int, selection *goquery.Selection) {
			attr, _ := selection.Attr("href")
			log.Println(attr)
			log.Println(selection.Text())
			//子收集器s
			//url := "https://wh.lianjia.com" + attr
			//sonC.Visit(url)
			//e.Request.Visit(url)
		})
	})

	////
	//sonC.OnRequest(func(r *colly.Request) {
	//	log.Println("sonC Visiting", r.URL)
	//})

	//c.OnHTML("", func(e *colly.HTMLElement) {
	//	log.Println(e)
	//})

	c.Visit(URL)
	c.Wait()
}
