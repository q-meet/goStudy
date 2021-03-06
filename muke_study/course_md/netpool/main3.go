package main

import (
	"fmt"
	colly "github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://hezhiqiang8909.gitbook.io/go/ch3-asm/ch3-06-func-again")

}