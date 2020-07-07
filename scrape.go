package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"github.com/gocolly/colly"
)

type quoteWord struct {
	Text string `json:"text"`
	Author string `json:"author"`
	Tags string `json:"tags"`

}

func writeFile(file []byte) {
	this := ioutil.WriteFile("file.json", file, 0644)
	if err := this; err != nil {
		panic(err)
	}
}
func serializeJSON(quote []quoteWord) {
	fmt.Println("Serializing ...")
	quotesJSON, _ := json.Marshal(quote)
	writeFile(quotesJSON)
	fmt.Println("Done Serializing √")
	fmt.Println(string(quotesJSON))
}

func main() {
	// Instantiate default collector
	c := colly.NewCollector()
	quotes := []quoteWord{}

	// On every a element which has href attribute call callback
	c.OnHTML("body > div > div.row.header-box", func(e *colly.HTMLElement) {

// Need to figure put how to write all the child texts

		quote := quoteWord{}
		quote.Text = e.ChildText("text")
		quote.Author = e.ChildText("author")
		quote.Tags = e.ChildText("tags")
		quotes = append(quotes, quote)

		// Print link
		fmt.Printf("Text: %q\n, Author: %q\n, Tags: %q\n", quote.Text, quote.Author, quote.Tags)

  // Before making a request print "Visiting ..."
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("http://quotes.toscrape.com/")
	serializeJSON(quotes)
}
