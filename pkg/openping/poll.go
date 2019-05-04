package ping

import (
	"log"
	"time"
)

// Poll polls, is a helper for our cmd client
func Poll(s Store, url string, locale LocationData) (uptime Uptime, latency Latency, metadata Metadata, sizes ContentSizes, err error) {
	request, err := GetRequest(url)
	if err != nil {
		return Uptime{
			Up:        0,
			Timestamp: time.Now(),
			RC:        0,
			URL:       url,
			Locale:    locale.Locale,
			Country:   locale.Country,
		}, Latency{}, Metadata{}, ContentSizes{}, err
	}
	uptime, latency, metadata, sizes, err = locale.GetDocument(request)
	log.Printf("%s %v %d bytes - %v", url, uptime.RC, metadata.Bytes, latency.TotalLatency)
	if err != nil {
		return Uptime{
			Up:        0,
			Timestamp: time.Now(),
			RC:        0,
			URL:       url,
			Locale:    locale.Locale,
			Country:   locale.Country,
		}, Latency{}, Metadata{}, ContentSizes{}, err
	}
	s.Update(uptime, latency, metadata, sizes)
	return uptime, latency, metadata, sizes, nil
}
