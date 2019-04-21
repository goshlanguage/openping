package ping

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetDocument fetches a document and ensures it gets a response.
// Assumes network connectivity from the test location
func TestGetDocument(t *testing.T) {
	doc, err := GetDocument("https://google.com")
	if err != nil {
		t.Errorf("Failed to fetch Google, err: %v", err.Error())
	}

	assert.NotEqual(t, doc, "", "Error, fetched empty doc for Google")
}
