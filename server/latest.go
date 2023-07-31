/**
 * Returns a map containing data regarding the latest version of the sdk
 *
 * Endpoint: /versions/latest/{channel} -> channel defaults to "stable" if not provided
 * Result: {
 *  body: {
 *	 channel: string
 *	 version: string,
 *	 date: string,
 *  }
 * }
 */
package server

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Yakiyo/dinx/utils"
	"github.com/gofiber/fiber/v2"
	json "github.com/json-iterator/go"
)

func init() {
	App.Get("versions/latest/:channel?", latestHandler)
}

func latestHandler(c *fiber.Ctx) error {
	channel := c.Params("channel", "stable")
	if !utils.IsValidChannel(channel) {
		return utils.Error(fmt.Sprintf("Argument `%v` is not a valid channel. Must be one of %v", channel, strings.Join(utils.Channels[:], ", ")), c)
	}
	url := fmt.Sprintf("https://storage.googleapis.com/dart-archive/channels/%v/release/latest/VERSION", channel)
	resp, err := http.Get(url)
	if err != nil {
		return utils.Error(err.Error(), c)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return utils.Error(err.Error(), c)
	}

	out := output{}

	err = json.Unmarshal(data, &out)

	if err != nil {
		return utils.Error(err.Error(), c)
	}

	return c.JSON(map[string]map[string]string{
		"body": {
			"channel": channel,
			"version": out.Version,
			"date":    out.Date,
		},
	})
}

type output struct {
	Version string `json:"version"`
	Date    string `json:"date"`
}
