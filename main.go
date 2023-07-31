package main

import (
	"log"
	"os"
	"strconv"

	"github.com/Yakiyo/dinx/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	if _, err := strconv.Atoi(port); err != nil {
		log.Fatalf("Value of %v for port is invalid. Must be an int", port)
	}

	server.App.Listen(":" + port)
}
