package main

import (
	"bytes"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
	"log"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector(colly.AllowURLRevisit())

	log.Println("---")
	//Rotate two socks5 proxies
	rp, err := proxy.RoundRobinProxySwitcher("http://127.0.0.1:1085", "http://127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	c.SetProxyFunc(rp)

	// Print the response
	c.OnResponse(func(r *colly.Response) {
		log.Printf("Proxy Address: %s\n", r.Request.ProxyURL)
		log.Printf("%s\n", bytes.Replace(r.Body, []byte("\n"), nil, -1))
	})

	c.OnError(func(response *colly.Response, err error) {
		log.Println("Error")
		log.Println(response)
		log.Println(err)
	})

	// Fetch httpbin.org/ip five times
	for i := 0; i < 2; i++ {
		c.Visit("https://www.baidu.com/")
	}
}
