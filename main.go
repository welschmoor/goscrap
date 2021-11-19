package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	var results = map[string]string{}
	urls := []string{
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://www.airbnb.com",
	}
	for _, each := range urls {
		result := "OK"
		err := VisitURL(each)
		if err != nil {
			result = "FAILED"
		}
		results[each] = result
	}
	for url, responseStatus := range results {
		fmt.Println(url, responseStatus)
	}
}

// this function visits URL using http package
func VisitURL(url string) {
	fmt.Println("<><>", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return errors.New("response failed")
	}
	return nil
}
