package main

import (
	"strings"
)

// The values received from the api are in the format of
// `channels/{{channel}}/release/{{version}}/`. This func
// strips the initial part and only gets the version
func stripVers(val string) string {
	subs := strings.Split(val, "/")
	return subs[len(subs)-2]
}
