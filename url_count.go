package main

import (
	"io"
	"log"
	"net/http"
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
func getBodyLen2(url string, ch chan struct{}) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error")
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	log.Printf("Getting body for %s", url)
	charsLen := utf8.RuneCountInString(string(body))
	log.Printf("chars: %d", charsLen)
	ch <- struct{}{}
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
		"https://www.adobe.com", "https://www.craigslist.com"}
	len := make(chan struct{})

	for _, url := range urls {
		go getBodyLen2(url, len)
	}
	for range urls {
		<-len
	}

}
func main() {
	GetBodyLens2()
}
