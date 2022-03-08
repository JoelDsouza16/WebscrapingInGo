package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gocolly/colly"
)

type InternshipDetails struct {
	Title              string   `json:"title"`
	Company            string   `json:"company"`
	Stipend            string   `json:"stipend"`
	StartDate          string   `json:"start_date"`
	EndDate            string   `json:"end_date"`
	Duration           string   `json:"duration"`
	Location           []string `json:"location"`
	ViewDescriptionURL string   `json:"view_description_URL"`
}

func ScrapeInternshala() []InternshipDetails {
	listOfInternship := []InternshipDetails{}

	location := []string{}
	c := colly.NewCollector()

	c.OnHTML("div#list_container", func(e *colly.HTMLElement) {

		e.ForEach("div.individual_internship", func(_ int, el *colly.HTMLElement) {
			singleJob := InternshipDetails{}
			singleJob.Title = el.ChildText("div.heading_4_5")
			singleJob.Company = el.ChildText("a.link_display_like_text")

			el.ForEach("a.location_link", func(_ int, e2 *colly.HTMLElement) {
				location = append(location, e2.Text)
			})
			singleJob.Location = location
			singleJob.Stipend = el.ChildText("span.stipend")

			listOfInternship = append(listOfInternship, singleJob)

			location = []string{}
		})
		// fmt.Println("Scrapping Complete")
	})

	c.Visit("https://internshala.com/internships")

	return listOfInternship
}

func main() {
	fmt.Println("Started the server...")
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/internshala", HelloInternshala)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server Is Up")
}
func HelloInternshala(w http.ResponseWriter, r *http.Request) {
	data := ScrapeInternshala()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
