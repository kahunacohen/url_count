package main

import (
	"io"
	"log"
	"net/http"
	"unicode/utf8"
)

var urls = []string{"ht://www.google.com", "https://www.walmart.com", "https://www.amazon.com", "https://www.nytimes.com",
	"https://www.trello.com", "https://mytzedakah.com/create-fund/1",
	"https://www.adobe.com", "https://wikipedia.org", "https://www.yahoo.com", "https://www.ncbi.nlm.nih.gov", "https://npr.org"}

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
func getBodyLen2(url string, charLengths chan ResponseChan) {
	resp, err := http.Get(url)
	if err != nil {
		charLengths <- ResponseChan{CharLength: 0, Error: err}
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		charLengths <- ResponseChan{CharLength: 0, Error: err}
		return
	}
	log.Printf("Getting body for %s", url)
	charLengths <- ResponseChan{CharLength: utf8.RuneCountInString(string(body)), Error: nil}
}

func GetBodyLens() {
	for _, url := range urls {
		len, _ := getBodyLen(url)
		log.Printf("characters: %d", len)
	}
}

type ResponseChan struct {
	CharLength int
	Error      error
}

func GetBodyLens2() {

	bodyLengths := make(chan ResponseChan, len(urls))
	for _, url := range urls {
		go getBodyLen2(url, bodyLengths)
	}
	for i, url := range urls {
		response := <-bodyLengths
		if response.Error != nil {
			log.Printf("can't get body length of %s: %v", url, response.Error)
		} else {
			log.Printf("Number of chars: %d", response.CharLength)
		}
		if i == len(urls)-1 {
			close(bodyLengths)
		}

	}

}
func main() {
	// GetBodyLens()
	GetBodyLens2()
}
