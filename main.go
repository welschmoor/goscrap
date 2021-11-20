package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type itemDict struct {
	id    string
	title string
	price string
}

//
var baseURL string = "https://www.ebay-kleinanzeigen.de/s-50321/fahrrad-pegasus/k0l1158r10"

// MAIN this is for all pages:
func main() {
	var jobs []itemDict
	totalPages := getPages()

	for i := 0; i < totalPages; i++ {
		listingsOnSinglePage := getPage(i)
		jobs = append(jobs, listingsOnSinglePage...)
	}
	// fmt.Println(jobs)
	for _, each := range jobs {

		fmt.Println(each.id, each.title, each.price)
	}
	saveItems(jobs)
}

// MAIN this one is for first page of results only:
// func main() {
// 	getPage(0) // only scrape the 1st page
// }

func saveItems(jobs []itemDict) {
	file, err := os.Create("items.csv") // creating a file
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"ID", "Title", "Price"}
	wErr := w.Write(headers)
	checkErr((wErr))

	for _, job := range jobs {
		itemSlice := []string{job.id, job.title, job.price}
		wErr := w.Write(itemSlice)
		checkErr(wErr)
	}
}

// pages on kleinanzeigen sind in der url mit /seite:1/ und /seite:2/ usw... codiert
func getPage(page int) []itemDict {
	var itemList []itemDict
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
		item1 := extractOneItem(s) // refactored funcction
		itemList = append(itemList, item1)
	})

	return itemList
}

func extractOneItem(s *goquery.Selection) itemDict {
	id, _ := s.Attr("data-adid") // finding id of each item
	title := cleanString(s.Find(".text-module-begin a").Text())
	price := cleanString(s.Find(".aditem-main--middle--price").Text())

	// fmt.Println(id, title, "Price: ", strings.TrimSpace(price))
	return itemDict{
		id:    id,
		title: title,
		price: price}

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

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
