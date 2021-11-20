package main

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

func main() {
	c := make(chan requestResult)
	results := make(map[string]string)
	urls := []string{
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://www.airbnb.com",
	}

	for _, url := range urls {
		go VisitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

// this function visits URL using http package
func VisitURL(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		c <- requestResult{url: url, status: "FAILED"}
	} else {
		c <- requestResult{url: url, status: "OK"}
	}
}
