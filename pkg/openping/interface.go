package ping

// Store is an interface for storing data
type Store interface {
	Update(uptime Uptime, latency Latency, meta Metadata, size ContentSizes) (err error)
}
