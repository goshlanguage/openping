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
	backend string
	poll    int
	sites   string
	store   ping.Store
	whc     ping.WebHookConfig
}

func main() {
	var store ping.Store
	var mongoURL, mongoUser, mongoPassword string
	var err error
	config := Config{}
	locale := ping.LocationData{}

	flag.StringVar(&config.backend, "backend", "", "The backend for storage, default is in memory.")
	flag.StringVar(&mongoURL, "mongodb-url", "", "The mongodb connection url to connect to your mongodb instance")
	flag.StringVar(&mongoUser, "mongodb-user", "", "The mongodb user to connect as")
	flag.StringVar(&mongoPassword, "mongodb-password", "", "The mongodb password to authenticate with")
	flag.IntVar(&config.poll, "poll-period", 60, "The poll period to wait between scrapes")
	flag.StringVar(&config.sites, "sites", "", "A common delimited string of sites to ping, ex: \"https://google.com,https://github.com\"")
	flag.StringVar(&config.whc.IconEmoji, "slack-icon", ":hankey:", "The emoji icon for your slack messages.")
	flag.StringVar(&config.whc.Username, "slack-username", "OpenPing Bot", "The display name for your slack messages.")
	flag.StringVar(&config.whc.WebhookURL, "slack-url", "", "Enter your slack webhook-url.")
	flag.StringVar(&locale.Locale, "locale", "", "If set, locale will store the location of the poller in the datastore")
	flag.StringVar(&locale.Country, "country", "", "If set, country will store the country of the poller in the datastore")
	flag.Parse()

	log.Print("Starting openping daemon")
	sites := strings.Split(config.sites, ",")

	if config.backend == "" {
		log.Printf("Creating in-memory document store")
		store = ping.NewDocumentStore()
	}

	// Here, if the NewMongoStore factory returns an error, we don't output mongoURL 
	// as to keep any user/pass out of stdout.
	if mongoURL != "" {
		log.Printf("Connecting to MongoDB: %v", mongoURL)
		if mongoUser != "" && mongoPassword != "" {
			log.Printf("Configured mongo auth with environment vars.")
			store, err = ping.NewMongoStore(mongoURL, mongoUser, mongoPassword)
			if err != nil {
				log.Printf("Error connecting to MongoDB, Error: %v", err.Error())
			}
		} else {
			store, err = ping.NewMongoStore(mongoURL, "", "")
			if err != nil {
				panic(fmt.Sprintf("Error connecting to MongoDB url: %v\tError: %v", mongoURL, err.Error()))
			}
		}
	}

	// Main loop, iterate through our sites and fetch them every n seconds
	for {
		for _, site := range sites {
			uptime, latency, _, _, err := ping.Poll(store, site, locale)
			if err != nil {
				if config.whc.WebhookURL != "" {
					message := fmt.Sprintf("Error detected for %s: %v", site, err.Error())
					config.whc.Alert(message)
				}
			}
			if uptime.Up == 0 {
				if config.whc.WebhookURL != "" {
					message := fmt.Sprintf("Site unhealthy, received a %v response code for: %s", uptime.RC, site)
					config.whc.Alert(message)
				}
			}
			if latency.TotalLatency > 3 {
				if config.whc.WebhookURL != "" {
					message := fmt.Sprintf("Latency alert, request to %s took %v seconds", site, latency.TotalLatency)
					config.whc.Alert(message)
				}
			}
		}
		// Sleep after the sites have been scraped
		time.Sleep(time.Duration(config.poll) * time.Second)
	}
}
