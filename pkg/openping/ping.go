package ping

import (
	"io/ioutil"
	"net/http"
	"time"
)

// GetRequest is a helper function that sets up the request. This is broken off into a help to improve testability.
func GetRequest(url string) (request *http.Request, err error) {
	// Create and modify HTTP request before sending
	request, err = http.NewRequest("GET", url, nil)
	return request, err
}

// GetDocument returns the HTML document to be stored in the document store for further analysis.
// This will be refactored to use channels
func GetDocument(request *http.Request) (document string, err error) {
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	request.Header.Set("User-Agent", "MobileSafari/604.1 CFNetwork/978.0.7 Darwin/18.5.0")

	// Make request
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)

	return string(data), nil
}
