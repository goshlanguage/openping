package ping

import "time"

// Status stores request metadata
type Status struct {
	timestamp time.Time
	state     string
}

// DocumentMemoryStore is an in memory store to track state of Documents.
// Documents are stored using the URL as the key, as is state.
type DocumentMemoryStore struct {
	Documents map[string][]string
	State     map[string]Status
}

// NewDocumentStore is a factory for DocumentMemoryStores
func NewDocumentStore() *DocumentMemoryStore {
	docs := make(map[string][]string)
	state := make(map[string]Status)
	return &DocumentMemoryStore{
		Documents: docs,
		State:     state,
	}
}

// Update takes a url and a document, and stores it in memory
func (dms *DocumentMemoryStore) Update(url string, rc int, latency time.Duration, document string) {
	// Check for nil map
	if dms.Documents[url] == nil {
		dms.Documents[url] = []string{}
	}
	dms.Documents[url] = append(dms.Documents[url], document)
}
