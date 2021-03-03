package main

import (
	"io"
	"log"
	"net/http"
	"unicode/utf8"
)

var urls = []string{"https://www.google.com", "https://www.walmart.com", "https://www.amazon.com", "https://www.nytimes.com",
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
func getBodyLen2(url string, charLengths chan int) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error")
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	log.Printf("Getting body for %s", url)
	charLengths <- utf8.RuneCountInString(string(body))
}

func GetBodyLens() {
	for _, url := range urls {
		len, _ := getBodyLen(url)
		log.Printf("characters: %d", len)
	}
}

func GetBodyLens2() {

	bodyLengths := make(chan int, len(urls))
	for _, url := range urls {
		go getBodyLen2(url, bodyLengths)
	}
	for i := range urls {
		v := <-bodyLengths
		log.Println(v)
		if i == len(urls)-1 {
			close(bodyLengths)
		}

	}

}
func main() {
	GetBodyLens()
	GetBodyLens2()
}
