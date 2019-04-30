package ping

import "time"

// Uptime stores uptime information.
type Uptime struct {
	Timestamp time.Time // The time of the request
	Up        bool      // A 200 Status code, free from common errors
	RC        int       // Return code
	URL       string
}

// Latency contains measurements of time for page loads
type Latency struct {
	DNSLookup    time.Duration // Time from DNSLookupStart to DNSLookupDone
	Timestamp    time.Time     // The time of the request
	TLSHandshake time.Duration // Time from TLSHandshakeStart to TLSHandshakeDone
	TotalLatency time.Duration // Time from connection to response
	TTFB         time.Duration // Time from Connection started to FirstByte Received
	URL          string
}

// ContentSizes breaks down the sizes of various parts of a request
type ContentSizes struct {
	CSSSize      int // Sum of Size in bytes of all stylesheets
	DocumentSize int // Sum of size in bytes of the fetched document
	FontSize     int // Sum of all loaded fonts
	ImageSize    int // Sum of Size in bytes of all jpg,jpeg,png,gif,svg,and ico links
	ScriptSize   int // Sum of all Javascript linked script files and inline scripts
	URL          string
}

// Metadata stores basic meta about a request
type Metadata struct {
	Bytes     int       // Size of the request
	Document  string    // A copy of the pulled document
	SHASum    string    // A SHA sum fingerprint of the document
	Timestamp time.Time // The time of the request
	URL       string
}
