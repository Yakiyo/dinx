package main

import (
	"io"
	"net/http"

	"github.com/Yakiyo/dinx/utils"
)

func fetch(channel string) {
	url := utils.MakeUrl(channel)
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	_ = parse(bytes)
}
