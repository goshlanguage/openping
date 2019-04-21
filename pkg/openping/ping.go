package ping

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// GetDocument returns the HTML document to be stored in the document store for further analysis.
// This will be refactored to use channels
func GetDocument(url string) (document string, err error) {
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// Create and modify HTTP request before sending
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Set("User-Agent", "MobileSafari/604.1 CFNetwork/978.0.7 Darwin/18.5.0")

	// Make request
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)

	return string(data), nil
}
