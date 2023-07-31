package server

import (
	"fmt"
	"strings"

	"github.com/Yakiyo/dinx/utils"
	"github.com/gofiber/fiber/v2"
)

func init() {
	App.Get("/versions", versionsHandler)
}

func versionsHandler(c *fiber.Ctx) error {
	channel := c.Query("channel", "stable")
	if !utils.IsValidChannel(channel) {
		return utils.Error(fmt.Sprintf("Argument %v of type channel is invalid. Must be one of %v", channel, strings.Join(utils.Channels[:], ", ")), c)
	}
	js, err := utils.ReadVersions("versions.json")

	if err != nil {
		return err
	}
	val := js[channel]
	return c.JSON(map[string][]string{
		channel: val,
	})
}
