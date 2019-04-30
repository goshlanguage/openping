package ping

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPoll(t *testing.T) {
	inmem := NewDocumentStore()
	server := httptest.NewServer(http.HandlerFunc(Mock200Handler))
	// Close the server when test finishes
	defer server.Close()

	// Test poll 1
	uptime, latency, metadata, _, err := Poll(inmem, server.URL)
	assert.True(t, latency.TotalLatency < (5*time.Second))
	assert.Equal(t, 200, uptime.RC)
	assert.True(t, metadata.Bytes > 0, "Got %v bytes of data, expected some amount of data instead.", metadata.Bytes)
	assert.NoError(t, err)

	// Test poll 2
	uptime, latency, metadata, _, err = Poll(inmem, server.URL)
	assert.True(t, latency.TotalLatency < (5*time.Second))
	assert.Equal(t, 200, uptime.RC)
	assert.NoError(t, err)
}