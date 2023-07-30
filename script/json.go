package main

import (
	"log"

	json "github.com/json-iterator/go"
)

type Resp struct {
	Prefixes []string `json:"prefixes"`
}

func parse(data []byte) Resp {
	resp := Resp{}

	err := json.Unmarshal(data, &resp)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	return resp
}
