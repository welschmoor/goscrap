package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

//
var baseURL string = "https://www.ebay-kleinanzeigen.de/s-50321/fahrrad-pegasus/k0l1158r10"

func main() {
	totalPages := getPages()

	for i := 0; i < totalPages; i++ {
		getPage(i)
	}
}

// pages on kleinanzeigen sind in der url mit /seite:1/ und /seite:2/ usw... codiert
func getPage(page int) {
	pageURL := "https://www.ebay-kleinanzeigen.de/s-50321/seite:" + strconv.Itoa(page+1) + "/fahrrad-pegasus/k0l1158r200"
	fmt.Println("pageURL:  ", pageURL)
}

// return number of pages
func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages += s.Find("a").Length() // how many a-tags are there
	})

	fmt.Println(doc)
	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err) // this kills the program
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Req failed with status: ", res.StatusCode)
	}
}
