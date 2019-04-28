package ping

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetDiffPercentage(t *testing.T){ 
	doc1 := `
	<html><head><title>Test</title></head><body><h1>OMG Great Test!</h1></body></html>
	`
	doc2 := `
	<html><head><title>Test</title></head><body><h1>OMG This test kind of sucks really</h1></body><script src="jQueryLOL"></script></html>
	`
	drift := GetDiffPercentage(doc1, doc2)
	t.assert.

}