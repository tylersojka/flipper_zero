package main

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/spf13/viper"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func sendsms() {
	vi := viper.New()
	vi.SetConfigFile("config.yaml")
	vi.ReadInConfig()

	twilio_sid := vi.GetString("twilio_sid")
	to_number := vi.GetString("to_number")
	from_number := vi.GetString("from_number")
	twilio_token := vi.GetString("twilio_token")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: twilio_sid,
		Password: twilio_token,
	})

	params := &api.CreateMessageParams{}
	params.SetTo(to_number)
	params.SetFrom(from_number)
	params.SetBody("GO BUY A FLIPPER ZERO")

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("SMS sent successfully!")
	}
}
func main() {

	scrapeUrl := "https://shop.flipperzero.one/password"

	c := colly.NewCollector(colly.AllowedDomains("shop.flipperzero.one", "www.shop.flipperzero.one"))

	// find and print the coming soon tag
	c.OnHTML("span.h2", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})

	// find the busy updating tag
	c.OnHTML("div.password-message", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})

	// print out the website visited and the status code
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL, r.StatusCode)
	})

	// handle and print out any errors
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit(scrapeUrl)
}
