package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

func main() {
	html := `
<div class="contentBottom clear"><div class="crumbs fl"><a href="/">武汉房产网</a><span>&nbsp;&gt;&nbsp;</span><a href="/ershoufang/">武汉二手房</a><span>&nbsp;&gt;&nbsp;</span><a href="/ershoufang/wuchang/">武昌、洪山二手房</a><span>&nbsp;&gt;&nbsp;</span><h1><a href="/ershoufang/chuhehanjie/">楚河汉街二手房</a></h1></div><div class="page-box fr"><div class="page-box house-lst-page-box" comp-module="page" page-url="/ershoufang/chuhehanjie/pg{page}" page-data="{&quot;totalPage&quot;:7,&quot;curPage&quot;:1}"><a class="on" href="/ershoufang/chuhehanjie/" data-page="1">1</a><a href="/ershoufang/chuhehanjie/pg2" data-page="2">2</a><a href="/ershoufang/chuhehanjie/pg3" data-page="3">3</a><span>...</span><a href="/ershoufang/chuhehanjie/pg7" data-page="7">7</a><a href="/ershoufang/chuhehanjie/pg2" data-page="2">下一页</a></div>
</div></div>
			`

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find(".page-box.house-lst-page-box>a:last-child").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Attr("href"))
	})
}
