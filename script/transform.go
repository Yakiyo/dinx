package main

import (
	"regexp"
	"strings"
)

// Simple regex to check string contains at least the main vers that is x.y.z
var re = regexp.MustCompile(`([1-9]+)\.([1-9]+)\.([1-9]+)`)

// The values received from the api are in the format of
// `channels/{{channel}}/release/{{version}}/`. This func
// strips the initial part and only gets the version
func stripVers(val string) string {
	subs := strings.Split(val, "/")
	return subs[len(subs)-2]
}
