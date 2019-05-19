package ping

import "time"

// Uptime stores uptime information.
type Uptime struct {
	Timestamp time.Time `json:"timestamp"` // The time of the request
	Up        int       `json:"up"`        // 1 if request returns a 200 Status code, free from common errors
	RC        int       `json:"rc"`        // Return code
	URL       string    `json:"url"`
	Locale    string    `json:"locale,omitempty"`
	Country   string    `json:"country,omitempty"`
}

// Latency contains measurements of time for page loads
type Latency struct {
	DNSLookup    float64   `json:"dns_lookup"`    // Time from DNSLookupStart to DNSLookupDone
	Timestamp    time.Time `json:"timestamp"`     // The time of the request
	TLSHandshake float64   `json:"tls_handshake"` // Time from TLSHandshakeStart to TLSHandshakeDone
	TotalLatency float64   `json:"total_latency"` // Time from connection to response
	TTFB         float64   `json:"ttfb"`          // Time from Connection started to FirstByte Received
	URL          string    `json:"url"`
	Locale       string    `json:"locale,omitempty"`
	Country      string    `json:"country,omitempty"`
}

// ContentSizes breaks down the sizes of various parts of a request
type ContentSizes struct {
	CSSSize      int    `json:"css_size"`      // Sum of Size in bytes of all stylesheets
	DocumentSize int    `json:"document_size"` // Sum of size in bytes of the fetched document
	FontSize     int    `json:"font_size"`     // Sum of all loaded fonts
	ImageSize    int    `json:"image_size"`    // Sum of Size in bytes of all jpg,jpeg,png,gif,svg,and ico links
	ScriptSize   int    `json:"script_size"`   // Sum of all Javascript linked script files and inline scripts
	URL          string `json:"url"`
	Locale       string `json:"locale,omitempty"`
	Country      string `json:"country,omitempty"`
}

// Metadata stores basic meta about a request
type Metadata struct {
	Bytes     int       `json:"bytes"`     // Size of the request
	Document  string    `json:"document"`  // A copy of the pulled document
	SHASum    string    `json:"shasum"`    // A SHA sum fingerprint of the document
	Timestamp time.Time `json:"timestamp"` // The time of the request
	URL       string    `json:"url"`
	Locale    string    `json:"locale,omitempty"`
	Country   string    `json:"country,omitempty"`
}

// LocationData should be used to track information about where requests are made from
type LocationData struct {
	Locale  string `json:"locale,omitempty"`
	Country string `json:"country,omitempty"`
	IP      string `json:"ip,omitempty"` // TODO write function to get IP
}
