package main

import (
	"flag"
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
}

func main() {
	config := Config{}
	flag.StringVar(&config.mongoDBURL, "mongodb-url", "", "The mongodb connection url to connect to your mongodb instance")
	flag.IntVar(&config.poll, "poll-period", 60, "The poll period to wait between scrapes")
	flag.StringVar(&config.sites, "sites", "", "A common delimited string of sites to ping, ex: \"https://google.com,https://github.com\"")
	flag.Parse()

	log.Print("Starting openping daemon")
	sites := strings.Split(config.sites, ",")
	for {
		for _, site := range sites {
			log.Printf("Fetching %s", site)
			doc, err := ping.GetDocument(site)
			if err != nil {
				log.Printf("Error fetching %s: %v", site, err.Error())
			}
			log.Printf("Size: %vk", len(doc)/1024)
		}
		// Sleep after the sites have been scraped
		time.Sleep(time.Duration(config.poll) * time.Second)
	}
}
