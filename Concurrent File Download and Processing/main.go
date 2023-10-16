package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

// Things about different type of channels
// Send Only channel ``
// Receive only channel
// One of the comman channel

type URLData struct {
	URL     string
	CONTENT []byte
	ERROR   error
}

func main() {
	URLs := []string{
		"https://www.soundguys.com/wp-content/uploads/2019/07/Screenshot-88-e1563568162263.jpg",
		"https://www.cnet.com/a/img/resize/6ba4173a19e1ca95f14a2ae302d8aa89773b873d/hub/2014/08/13/0c1cdd30-f845-4db4-a683-5e7637766125/pioneer-atmos-speakers.jpg?auto=webp&width=1200",
	}

	downloadChan := make(chan URLData)
	processChan := make(chan URLData)
	resultChan := make(chan URLData)
	var wg sync.WaitGroup
	for _, url := range URLs {
		wg.Add(1)
		go downloadFile(url, downloadChan, &wg)
	}
	for i := 0; i <= len(URLs); i++ {
		wg.Add(1)
		go ProcessAndSave(<-downloadChan, processChan, &wg)
	}
	wg.Wait()
	for i := 0; i < len(URLs); i++ {
		wg.Add(1)
		go ProcessAndSave(<-processChan, resultChan, &wg)
	}
	close(downloadChan)
	close(processChan)

	for i := 0; i < len(URLs); i++ {
		result := <-resultChan
		if result.ERROR != nil {
			fmt.Printf("Error processing %s: %v\n", result.URL, result.ERROR)
		} else {
			fmt.Printf("Processed and saved %s\n", result.URL)
		}
	}

	// Measure and display the time taken to complete the process
	elapsed := time.Since(time.Now())
	fmt.Printf("Total time taken: %s\n", elapsed)
}

func downloadFile(url string, ch chan URLData, wg *sync.WaitGroup) {
	defer wg.Done()
	res, err := http.Get(url)
	if err != nil {
		ch <- URLData{
			URL:     url,
			CONTENT: nil,
			ERROR:   err,
		}
		return
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		ch <- URLData{URL: url, ERROR: err}
		return
	}
	ch <- URLData{URL: url, CONTENT: data, ERROR: nil}
}

func ProcessAndSave(data URLData, resultCh chan<- URLData, wg *sync.WaitGroup) {
	defer wg.Done()

	fileName := "output_" + data.URL[8:] // Use part of the URL as the file name

	err := os.WriteFile(fileName, data.CONTENT, 0666)
	if err != nil {
		data.ERROR = err
	}

	resultCh <- data
}
