package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {

	fmt.Println("Welcome to a web scraping project demo..")
	location := ""
	c := colly.NewCollector()

	// Selection of Element - "div",  ID - "list_container"
	c.OnHTML("div#list_container", func(e *colly.HTMLElement) {

		// Selection of Element - "div",  css - "individual_internship"
		e.ForEach("div.individual_internship", func(_ int, el *colly.HTMLElement) {

			// extract the child text of (Element - "div",  css - "heading_4_5") 	& 	(Element - "a",  css - "link_display_like_text")
			fmt.Printf(" Title :: %v \n\t Company :: %v \n", el.ChildText("div.heading_4_5"), el.ChildText("a.link_display_like_text"))

			// Select (Element - "a",  css - "location_link") to format multiple locations
			el.ForEach("a.location_link", func(_ int, e2 *colly.HTMLElement) {
				location = location + e2.Text + ","
			})
			fmt.Printf("\t\t Location - %v", location[:len(location)-1])
			fmt.Printf("\n \tStipend - %v \n", el.ChildText("span.stipend"))
			location = ""
		})
		fmt.Println("Scrapping Complete")
	})
	c.Visit("https://internshala.com/internships")
	// time.Sleep(5 * time.Second)
}
