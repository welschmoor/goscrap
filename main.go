package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedItem struct {
	id    string
	title string
	price string
}

//
var baseURL string = "https://www.ebay-kleinanzeigen.de/s-50321/fahrrad-pegasus/k0l1158r10"

func main() {
	// totalPages := getPages()

	// for i := 0; i < totalPages; i++ {
	// 	getPage(i)
	// }
	getPage(0)
}

// pages on kleinanzeigen sind in der url mit /seite:1/ und /seite:2/ usw... codiert
func getPage(page int) {
	pageURL := "https://www.ebay-kleinanzeigen.de/s-50321/seite:" + strconv.Itoa(page+1) + "/fahrrad-pegasus/k0l1158r200"
	fmt.Println("pageURL:  ", pageURL)

	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	seachItems := doc.Find(".aditem") // .aditem is the maindiv of each posting
	seachItems.Each(func(i int, s *goquery.Selection) {
		id, _ := s.Attr("data-adid") // finding id of each item
		fmt.Println(" ")

		title := s.Find(".text-module-begin a").Text()
	
		price := s.Find(".aditem-main--middle--price").Text()
		fmt.Println(id, title, "Price: ", strings.TrimSpace(price))
		// extractedItem{id:id, title:title, price:price}
	})
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

// func cleanString(str string) string {

// }
