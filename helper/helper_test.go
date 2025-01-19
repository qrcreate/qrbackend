package helper

import (
	"testing"
)

func TestGetPresensiThisMonth(t *testing.T) {
	uri := SRVLookup("mongodb+srv://xx:xxx@cxxx.xxx.mongodb.net/")
	print(uri)
}

func TestSRVLookup(t *testing.T) {
	uri := SRVLookup("mongodb+srv://xx:xxx@cxxx.xxx.mongodb.net/")
	if uri == "" {
		t.Errorf("expected non-empty URI, got '%s'", uri)
	}
}
