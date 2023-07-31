package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	json "github.com/json-iterator/go"
)

var App *fiber.App = fiber.New(fiber.Config{
	AppName: "dinx",
	JSONEncoder: json.Marshal,
	JSONDecoder: json.Unmarshal,
})

func init() {
	App.Use(recover.New())

	App.Get("/", func(c *fiber.Ctx) error {
		c.SendString("Hello World :)")
		return nil
	})
}
