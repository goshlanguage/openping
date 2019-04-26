package ping

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptrace"
	"time"
)

// GetRequest is a helper function that sets up the request. This is broken off into a help to improve testability.
func GetRequest(url string) (request *http.Request, err error) {
	return http.NewRequest("GET", url, nil)
}

// GetDocument returns the HTML document to be stored in the document store for further analysis.
// This will be refactored to use channels
func GetDocument(request *http.Request) (document string, rc int, latency time.Duration, err error) {
	var dns0, dns1, tls0, tls1, ttfb0, ttfb1, conn1 time.Time
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	request.Header.Set("User-Agent", "MobileSafari/604.1 CFNetwork/978.0.7 Darwin/18.5.0")

	trace := &httptrace.ClientTrace{
		ConnectStart: func(string, string) {
			ttfb0 = time.Now()
		},
		GotFirstResponseByte: func() {
			ttfb1 = time.Now()
		},
		DNSStart: func(dnsInfo httptrace.DNSStartInfo) {
			dns0 = time.Now()
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			dns1 = time.Now()
		},
		TLSHandshakeStart: func() {
			tls0 = time.Now()
		},
		TLSHandshakeDone: func(tls.ConnectionState, error) {
			tls1 = time.Now()
		},
	}

	request = request.WithContext(httptrace.WithClientTrace(request.Context(), trace))
	if _, err := http.DefaultTransport.RoundTrip(request); err != nil {
		log.Printf("Oops, %v", err.Error())
	}

	response, err := client.Do(request)
	conn1 = time.Now()

	dnsLookupTime := dns1.Sub(dns0)
	log.Printf("DNS Time info: %v", dnsLookupTime)

	tlsHandshakeTime := tls1.Sub(tls0)
	log.Printf("TLS Handshake time: %v", tlsHandshakeTime)

	ttfb := ttfb1.Sub(ttfb0)
	log.Printf("TTFB time: %v", ttfb)

	latency = conn1.Sub(ttfb0)
	log.Printf("All done: %v", latency)

	if err != nil {
		return "", 0, latency, err
	}
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)

	return string(data), response.StatusCode, latency, nil
}
