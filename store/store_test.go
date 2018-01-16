package store

import (
	"testing"
)

func TestConnURI(t *testing.T) {
	const expectedURI = "postgres://:@:/?sslmode=disable"
	URI := connURI()

	if URI != expectedURI {
		t.Errorf("expected connURI to return %s, got: %s", expectedURI, URI)
	}
}
