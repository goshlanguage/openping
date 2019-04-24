package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	ping "github.com/ryanhartje/openping/pkg/openping"
)

// Config stores all the config info for openping.
type Config struct {
	backend    string
	mongoDBURL string
	poll       int
	sites      string
	whc        ping.WebHookConfig
}

func main() {
	var store ping.Store
	config := Config{}

	flag.StringVar(&config.backend, "backend", "", "The backend for storage, default is in memory.")
	flag.StringVar(&config.mongoDBURL, "mongodb-url", "", "The mongodb connection url to connect to your mongodb instance")
	flag.IntVar(&config.poll, "poll-period", 60, "The poll period to wait between scrapes")
	flag.StringVar(&config.sites, "sites", "", "A common delimited string of sites to ping, ex: \"https://google.com,https://github.com\"")
	flag.StringVar(&config.whc.IconEmoji, "slack-icon", ":hankey:", "The emoji icon for your slack messages.")
	flag.StringVar(&config.whc.Username, "slack-username", "OpenPing Bot", "The display name for your slack messages.")
	flag.StringVar(&config.whc.WebhookURL, "slack-url", "", "Enter your slack webhook-url.")

	flag.Parse()

	log.Print("Starting openping daemon")
	sites := strings.Split(config.sites, ",")

	if config.backend == "" {
		store = ping.NewDocumentStore()
	}
	// Main loop, iterate through our sites and fetch them every n seconds
	for {
		for _, site := range sites {
			rc, err := ping.Poll(store, site)
			if err != nil {
				if config.whc.WebhookURL != "" {
					message := fmt.Sprintf("Error detected for %s: %v", site, err.Error())
					config.whc.Alert(message)
				}
			}
			if rc != 200 {
				if config.whc.WebhookURL != "" {
					message := fmt.Sprintf("Site unhealthy, received a %s response code for: %s", string(rc), site)
					config.whc.Alert(message)
				}
			}
		}
		// Sleep after the sites have been scraped
		time.Sleep(time.Duration(config.poll) * time.Second)
	}
}
