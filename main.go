package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/getlantern/systray"
	"github.com/robfig/cron/v3"
	"github.com/soryuu/dogetray-go/icon"
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
	s.Cron.AddFunc("@every 10m", s.updatePrice)
	s.Cron.Start()
}

func (s *state) onExit() {
	s.Cron.Stop()
}

func (s *state) updatePrice() {
	res, err := http.Get("https://coinmarketcap.com/currencies/dogecoin/")

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	price := doc.Find(".priceValue").Text()
	fmt.Println("price")
	fmt.Println(price)
	floatPrice, err := strconv.ParseFloat(price[1:], 64)

	if err != nil {
		log.Fatal(err)
	}

	title := strconv.FormatFloat(math.Round(floatPrice*100)/100, 'f', -1, 64)

	systray.SetTitle("$" + title)
}
