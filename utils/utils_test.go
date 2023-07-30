package utils

import "testing"

func TestUrl(t *testing.T) {
	expected := `https://www.googleapis.com/storage/v1/b/dart-archive/o?delimiter=%2F&prefix=channels%2Fstable%2Frelease%2F&alt=json`
	out := MakeUrl("stable")
	if out != expected {
		t.Fatalf("Output did not match expected string\nReceived: %v\nExpected: %v", out, expected)
	}
}

func TestValidChannel(t *testing.T) {
	mock := map[string]bool{
		"stable": true,
		"dev":    true,
		"beta":   true,
		"blah":   false,
		"bruh":   false,
	}
	for k, v := range mock {
		res := IsValidChannel(k)
		if res != v {
			t.Fatalf("Failed to determine valid channel for %v. Expected %v, got %v", k, v, res)
		}
	}
}
