package ping

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestGetDocument fetches a document and ensures it gets a response.
// This test creates its own HTTP endpoint and serves it over TCP to set the test up and handle teardown.
// In this test, we insert the return of GetDocument 2 times into the in memory datastore to ensure we're
// storing multiple entries in memory.
func TestGetDocument(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(Mock200Handler))
	// Close the server when test finishes
	defer server.Close()
	req, err := GetRequest(server.URL)
	if err != nil {
		t.Fatalf(err.Error())
	}

	ld := LocationData{}

	uptime, latency, meta, _, err := ld.GetDocument(req)
	assert.True(t, latency.TotalLatency < 1, "Latency reading seems inaccurate")
	assert.Equal(t, meta.Document, "OK")
	assert.Equal(t, uptime.RC, 200)

	// Get a second request, ensure against time bug regression
	req, err = GetRequest(server.URL)
	if err != nil {
		t.Fatalf(err.Error())
	}

	uptime, latency, meta, _, err = ld.GetDocument(req)
	assert.True(t, latency.TotalLatency < 1, "Latency reading seems inaccurate, got: %v", latency.TotalLatency)
}

// TestUnresolved tests the response from an unresolvable URL.
// We do this to ensure an unresolvable entry will not make the service panic, and that it handles the failure
// gracefully
func TestUnresolved(t *testing.T) {
	ld := LocationData{}
	req, err := GetRequest("http://foo.bar")
	if err != nil {
		t.Errorf("Error forming request for foo.bar: %v", err.Error())
	}
	uptime, _, _, _, err := ld.GetDocument(req)
	assert.EqualError(t, err, "dial tcp: lookup foo.bar: no such host")
	assert.Equal(t, 0, uptime.RC, "Return code for a timeout should be 0, the nil value.")
}

// Test400Response makes sure we handle a 400 response code correctly
func Test400Response(t *testing.T) {
	ld := LocationData{}
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(Mock400Handler))
	// Close the server when test finishes
	defer server.Close()
	req, err := GetRequest(server.URL)
	if err != nil {
		t.Fatalf(err.Error())
	}

	uptime, latency, meta, _, err := ld.GetDocument(req)
	assert.True(t, latency.TotalLatency < 1, "Latency seems inaccurate")
	assert.Equal(t, meta.Document, "Bad Request")
	assert.Equal(t, uptime.RC, 400)
}

func TestLatency(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(MockSleepyHandler))
	// Close the server when test finishes
	defer server.Close()
	req, err := GetRequest(server.URL)
	if err != nil {
		t.Fatalf(err.Error())
	}

	ld := LocationData{}
	uptime, latency, meta, _, err := ld.GetDocument(req)
	assert.True(t, latency.TotalLatency > 2, "Latency measurement seems inaccurate.")
	assert.Equal(t, meta.Document, "Slow Request")
	assert.Equal(t, uptime.RC, 200)
}

/* Deprecated until variable timeout is implemented
func TestTimeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(MockTimeoutHandler))
	// Close the server when test finishes
	defer server.Close()
	req, err := GetRequest(server.URL)
	if err != nil {
		t.Fatalf(err.Error())
	}

	uptime, latency, meta, _, err := GetDocument(req)
	assert.Equal(t, latency.TotalLatency, (0 * time.Second), "Latency measurement seems inaccurate. Expected 5+ seconds, got: %v", latency)
	assert.Equalf(t, "", meta.Document, "A timed out document should produce an empty document, got: %v", meta.Document)
	assert.Equalf(t, 0, uptime.RC, "A timed out request should have a 0 return code, got: %v", uptime.RC)
	assert.Error(t, err, "Expected an error for timing out, got nil instead.")
}
*/

func Mock200Handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "OK")
}

func Mock301Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMovedPermanently)
	io.WriteString(w, "301 in the house")
}

func Mock400Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	io.WriteString(w, "Bad Request")
}

func MockSleepyHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)
	io.WriteString(w, "Slow Request")
}

func MockTimeoutHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(6 * time.Second)
	io.WriteString(w, "TimedOut")
}
