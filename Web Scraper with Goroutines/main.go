package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type webData struct {
	URL   string
	TITLE string
}

// Challenge: Web Scraper with Goroutines
func main() {
	website := []string{
		"https://go.dev",
		"https://google.com",
		"https://amazone.com",
	}
	ch := make(chan webData, len(website))
	errorCh := make(chan error)
	reteLimite := time.Tick(500 * time.Millisecond)
	for _, url := range website {
		<-reteLimite
		go extractInfoFromURL(url, ch, errorCh)
	}
	for range website {
		data := <-ch
		fmt.Println("Title of the website", data.TITLE)
	}
}

func extractInfoFromURL(url string, ch chan webData, errCh chan error) {
	// Make an HTTP GET request to the specified URL.
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	res, err := client.Get(url)
	if err != nil {
		log.Printf("Error fetching %s: %v\n", url, err)
		ch <- webData{
			URL:   url,
			TITLE: err.Error()}
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Printf("Received non-200 status code for %s: %d\n", url, res.StatusCode)
		ch <- webData{URL: url, TITLE: "Error"}
		return
	}

	if res.StatusCode != 200 {
		log.Printf("status code is not 200")
		ch <- webData{
			URL:   url,
			TITLE: "error status code",
		}
		return
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Printf("doc error")
		ch <- webData{
			URL:   url,
			TITLE: "error status code",
		}
		return
	}
	title := doc.Find("title").Text()
	ch <- webData{
		URL:   url,
		TITLE: title,
	}
	// You can extract other information here, such as headings or specific data
}
