/**
 * Returns an array of all versions of the dart sdk
 *
 * Endpoint: /versions/all
 * Result: { body: []string }
 */

package server

import (
	"github.com/Yakiyo/dinx/utils"
	"github.com/gofiber/fiber/v2"
)

func init() {
	App.Get("/versions/all", allVersionsHandler)
}

func allVersionsHandler(c *fiber.Ctx) error {
	versions := []string{}
	js, err := utils.ReadVersions("versions.json")
	if err != nil {
		return utils.Error(err.Error(), c)
	}

	for _, v := range js {
		versions = append(versions, v...)
	}
	return c.JSON(map[string][]string{
		"body": versions,
	})
}
