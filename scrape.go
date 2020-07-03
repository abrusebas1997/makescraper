package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/gocolly/colly"
)

type recipe struct {
	Image string `json:"image"`
	Title string `json:"title"`
	Description int `json:"rating"`
	Ingredients string `json:"ingredients"`
	Instructions string `json:"preparation"`
	Equipment string `json:"Equipment"`
}

func writeFile(file []byte) {
	this := ioutil.WriteFile("file.json", file, 0644)
	if err := this; err != nil {
		panic(err)
	}
}
func serializeJSON(recipe []recipe) {
	fmt.Println("Serializing ...")
	recipeJSON, _ := json.Marshal(recipe)
	writeFile(recipeJSON)
	fmt.Println("Done Serializing √")
	fmt.Println(string(recipeJSON))
}

func main() {
	// Instantiate default collector
	c := colly.NewCollector()
	recipes := []recipe{}
	// On every a element which has href attribute call callback
	c.OnHTML(".row ol > li", func(e *colly.HTMLElement) {

// Need to figure put how to write all the child texts

		// recipe := recipe{}
		// recipe.Image = e.Attr("Image")
		// recipe.Title = e.ChildText("Apple pie")
		// recipe.Description = e.ChildText("Description")
		// recipe.Ingredients = e.ChildText("Ingredients")
		// recipe.Instructions = e.ChildText("Instructions")
		// recipe.Equipment = e.ChildText("Equipment")
		// recipes = append(recipes, recipe)

		// Print link
		fmt.Printf("Image: %q\n, Title: %q\n, Description: %q\n, Ingredients: %q\n, Instructions: %q\n, Equipment: %q\n", recipe.Image, recipe.Title, recipe.Description, recipe.Ingredients, recipe.Instructions, recipe.Equipment)
				// fmt.Printf("Link found: %q -> %s\n", e.Text, link)
	})
	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://pinchofyum.com/5-ingredient-marinated-tomatoes")
	serializeJSON(recipes)
}
