package main

import (
	"io"
	"log"
	"net/http"
	"sync"
	"unicode/utf8"
)

func getBodyLen(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	log.Printf("Getting body for %s", url)
	return utf8.RuneCountInString(string(body)), nil
}
func getBodyLen2(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error")
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	log.Printf("Getting body for %s", url)
	charsLen := utf8.RuneCountInString(string(body))
	log.Printf("chars: %d", charsLen)

}

func GetBodyLens() {
	urls := []string{"https://www.google.com", "https://www.nytimes.com",
		"https://www.trello.com", "https://mytzedakah.com/create-fund/1",
		"https://www.adobe.com", "https://www.craigslist.com"}
	for _, url := range urls {
		len, _ := getBodyLen(url)
		log.Printf("characters: %d", len)
	}
}

func GetBodyLens2() {

	urls := []string{"https://www.google.com", "https://www.nytimes.com",
		"https://www.trello.com", "https://mytzedakah.com/create-fund/1",
		"https://www.adobe.com", "https://wikipedia.org", "https://www.yahoo.com", "https://www.ncbi.nlm.nih.gov"}
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go getBodyLen2(url, &wg)
	}
	wg.Wait()

}
func main() {
	GetBodyLens2()
}
