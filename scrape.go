package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type food struct {
	Name string `json:"name"`
	Price int `json:"price"`
	Rating string `json:"rating"`
	Image string `json:"image"`
}

func writeFile(file []byte) {
	this := ioutil.WriteFile("file.json", file, 0644)
	if err := this; err != nil {
		panic(err)
	}
}
// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
                link := e.Attr("href")

		// Print link
                fmt.Printf("Link found: %q -> %s\n", e.Text, link)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://hackerspaces.org/")
}
