package ping

import "time"

// Store is an interface for storing data
type Store interface {
	Update(url string, rc int, latency time.Duration, document string)
}
