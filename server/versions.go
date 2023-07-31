/**
 * Returns an array of versions of the requested channel the dart sdk
 *
 * Endpoint: /versions/{channel}
 * Result: { body: []string }
 */

package server

import (
	"fmt"
	"strings"

	"github.com/Yakiyo/dinx/utils"
	"github.com/gofiber/fiber/v2"
)

func init() {
	App.Get("/versions/:channel", versionsHandler)
}

func versionsHandler(c *fiber.Ctx) error {
	channel := c.Params("channel")
	if !utils.IsValidChannel(channel) {
		return utils.Error(fmt.Sprintf("Argument `%v` is not a valid channel. Must be one of %v", channel, strings.Join(utils.Channels[:], ", ")), c)
	}
	js, err := utils.ReadVersions("versions.json")

	if err != nil {
		return err
	}
	val := js[channel]
	return c.JSON(map[string][]string{
		"body": val,
	})
}
