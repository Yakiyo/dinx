package main

import "github.com/Yakiyo/dinx/server"

func main() {
	server.App.Listen(":3000")
}
