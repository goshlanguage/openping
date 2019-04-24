package ping

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

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

	doc, rc, err := GetDocument(req)
	assert.Equal(t, doc, "OK")
	assert.Equal(t, rc, 200)
}

// TestUnresolved tests the response from an unresolvable URL
func TestUnresolved(t *testing.T) {
	req, err := GetRequest("http://foo.bar")
	if err != nil {
		t.Errorf("Error forming request for foo.bar: %v", err.Error())
	}
	_, _, err = GetDocument(req)
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

	doc, rc, err := GetDocument(req)
	assert.Equal(t, doc, "Bad Request")
	assert.Equal(t, rc, 400)
}

func Mock200Handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "OK")
}

func Mock400Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	io.WriteString(w, "Bad Request")
}
