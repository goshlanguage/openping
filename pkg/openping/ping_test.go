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
	server := httptest.NewServer(http.HandlerFunc(MockHandler))
	// Close the server when test finishes
	defer server.Close()
	req, err := GetRequest(server.URL)
	if err != nil {
		t.Fatalf(err.Error())
	}

	doc, err := GetDocument(req)
	assert.Equal(t, doc, "OK")
}

// TestUnresolved tests the response from an unresolvable URL
func TestUnresolved(t *testing.T) {
	req, err := GetRequest("http://foo.bar")
	if err != nil {
		t.Errorf("Error forming request for foo.bar: %v", err.Error())
	}
	_, err = GetDocument(req)
	assert.EqualError(t, err, "Get http://foo.bar: dial tcp: lookup foo.bar: no such host")
}

func MockHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "OK")
}
