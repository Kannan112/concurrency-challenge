package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// Challenge: Web Scraper with Goroutines
func main() {
	website := []string{
		"https://go.dev",
		"https://google.com",
		"amazone.com",
	}
	ch := make(chan interface{})
	errorCh := make(chan error)
	for _, url := range website {
		go extractInfoFromURL(url, ch, errorCh)
	}
	<-ch

}

func extractInfoFromURL(url string, ch chan interface{}, errCh chan error) error {
	// Make an HTTP GET request to the specified URL.
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		errCh := fmt.Errorf("HTTP request failed with status code: %d", response.StatusCode)
		<-errCh
	}

	// Parse the HTML content using goquery.
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return err
	}

	// Extract and print the title of the webpage.
	title := doc.Find("title").Text()
	fmt.Println("Title:", title)

	// You can extract other information here, such as headings or specific data.

	return nil
}
