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

func GetBodyLens() {
	urls := []string{"https://www.google.com", "https://www.nytimes.com",
		"https://www.trello.com", "https://mytzedakah.com/create-fund/1",
		"https://www.adobe.com"}

	for _, url := range urls {
		len, _ := getBodyLen(url)
		log.Printf("characters: %d", len)
	}
}
func main() {
	GetBodyLens()

}
