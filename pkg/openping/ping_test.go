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
// Assumes network connectivity from the test location
func TestGetDocument(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(Mock200Handler))
	// Close the server when test finishes
	defer server.Close()
	req, err := GetRequest(server.URL)
	if err != nil {
		t.Fatalf(err.Error())
	}

	doc, rc, latency, err := GetDocument(req)
	assert.True(t, latency < time.Second, "Latency reading seems inaccurate")
	assert.Equal(t, doc, "OK")
	assert.Equal(t, rc, 200)
}

// TestUnresolved tests the response from an unresolvable URL
func TestUnresolved(t *testing.T) {
	req, err := GetRequest("http://foo.bar")
	if err != nil {
		t.Errorf("Error forming request for foo.bar: %v", err.Error())
	}
	_, _, _, err = GetDocument(req)
	assert.EqualError(t, err, "Get http://foo.bar: dial tcp: lookup foo.bar: no such host")
}

// Test400Response makes sure we handle a 400 response code correctly
func Test400Response(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(Mock400Handler))
	// Close the server when test finishes
	defer server.Close()
	req, err := GetRequest(server.URL)
	if err != nil {
		t.Fatalf(err.Error())
	}

	doc, rc, latency, err := GetDocument(req)
	assert.True(t, latency < time.Second, "Latency seems inaccurate")
	assert.Equal(t, doc, "Bad Request")
	assert.Equal(t, rc, 400)
}

func TestLatency(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(MockSleepyHandler))
	// Close the server when test finishes
	defer server.Close()
	req, err := GetRequest(server.URL)
	if err != nil {
		t.Fatalf(err.Error())
	}

	doc, rc, latency, err := GetDocument(req)
	assert.True(t, latency > (2*time.Second), "Latency measurement seems inaccurate.")
	assert.Equal(t, doc, "Slow Request")
	assert.Equal(t, rc, 200)
}

func TestTimeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(MockTimeoutHandler))
	// Close the server when test finishes
	defer server.Close()
	req, err := GetRequest(server.URL)
	if err != nil {
		t.Fatalf(err.Error())
	}

	doc, rc, latency, err := GetDocument(req)
	assert.True(t, latency > (5*time.Second), "Latency measurement seems inaccurate.")
	assert.Equal(t, doc, "Slow Request")
	assert.Equal(t, rc, 200)
}

func Mock200Handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "OK")
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
	io.WriteString(w, "Slow Request")
}
