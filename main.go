package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/getlantern/systray"
	"github.com/robfig/cron/v3"
	"github.com/soryuu/doge-systray/icon"
)

type state struct {
	Price string
	Cron  *cron.Cron
}

func main() {
	s := &state{}

	fmt.Println("service started")
	systray.Run(s.onReady, s.onExit)
}

func (s *state) onReady() {
	systray.SetIcon(icon.Data)
	s.updatePrice()

	s.Cron = cron.New()
	s.Cron.AddFunc("@every 20s", s.updatePrice)
	s.Cron.Start()
}

func (s *state) onExit() {
	s.Cron.Stop()
}

func (s *state) updatePrice() {
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
