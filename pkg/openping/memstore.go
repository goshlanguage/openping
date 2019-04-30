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
	Uptimes   []Uptime
	Latencies []Latency
	Metas     []Metadata
	Sizes     []ContentSizes
}

// NewDocumentStore is a factory for DocumentMemoryStores
func NewDocumentStore() *DocumentMemoryStore {
	return &DocumentMemoryStore{
		Uptimes:   []Uptime{},
		Latencies: []Latency{},
		Metas:     []Metadata{},
		Sizes:     []ContentSizes{},
	}
}

// Update takes a url and a document, and stores it in memory
func (dms *DocumentMemoryStore) Update(uptime Uptime, latency Latency, meta Metadata, size ContentSizes) (err error) {
	dms.Uptimes = append(dms.Uptimes, uptime)
	dms.Latencies = append(dms.Latencies, latency)
	dms.Metas = append(dms.Metas, meta)
	dms.Sizes = append(dms.Sizes, size)
	return nil
}
