package ping

import (
	"time"

	_ "go.mongodb.org/mongo-driver"
)

// MongoStore contains information required to establish a mondodb connection and read/write data
type MongoStore struct {
	MongoDBURL string
}

// Update stores a document in a Mongo Document Store
func (ms *MongoStore) Update(url string, rc int, latency time.Duration, document string) {

}
