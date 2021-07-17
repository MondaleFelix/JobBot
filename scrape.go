package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Post struct {
	title    string `json:"title"`
	company  string
	location string
	link     string
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/

func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("#mosaic-provider-jobcards", func(e *colly.HTMLElement) {

		e.ForEach("a[data-jk]", func(_ int, e *colly.HTMLElement) {
			link := e.Attr("href")
			fmt.Println(link)

		})

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://www.indeed.com/jobs?q=ios+developer&l=San+Francisco")
}
