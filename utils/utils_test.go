package utils

import "testing"

func TestUrl(t *testing.T) {
	expected := `https://www.googleapis.com/storage/v1/b/dart-archive/o?delimiter=%2F&prefix=channels%2Fstable%2Frelease%2F&alt=json`
	out := MakeUrl("stable")
	if out != expected {
		t.Fatalf("Output did not match expected string\nReceived: %v\nExpected: %v", out, expected)
	}
}
