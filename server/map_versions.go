/**
 * Returns an map of channel and versions of the dart sdk
 *
 * Endpoint: /versions/map
 * Result: { body: map[string][]string }
 */

package server

import (
	"github.com/Yakiyo/dinx/utils"
	"github.com/gofiber/fiber/v2"
)

func init() {
	App.Get("/versions/map", mapVersionsHandler)
}

func mapVersionsHandler(c *fiber.Ctx) error {
	js, err := utils.ReadVersions("versions.json")
	if err != nil {
		return utils.Error(err.Error(), c)
	}

	return c.JSON(js)
}
