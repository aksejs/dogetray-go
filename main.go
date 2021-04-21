package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/getlantern/systray"
	"github.com/soryuu/doge-systray/icon"
)

func main() {
	fmt.Println("service started")
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(icon.Data)
	getPrice()
	systray.AddMenuItem("Quit", "Quit the whole app")
}

func onExit() {
	// clean up here
}
func getPrice() {
	// Request the HTML page.
	res, err := http.Get("https://coinmarketcap.com/currencies/dogecoin/")

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	price := doc.Find(".priceValue___11gHJ").Text()
	systray.SetTitle(price)
}
