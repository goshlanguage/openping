package ping

// Store is an interface for storing data
type Store interface {
	Update(url string, document string)
}
