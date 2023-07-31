package utils

import (
	"os"

	"github.com/gofiber/fiber/v2"
	json "github.com/json-iterator/go"
)

// Read json from file
func ReadVersions(path string) (map[string][]string, error) {
	f, err := os.ReadFile(path)
	m := map[string][]string{}

	if err != nil {
		return m, err
	}
	err = json.Unmarshal(f, &m)
	if err != nil {
		return m, err
	}
	return m, nil
}

// Convert a string to a { error: `msg` } map
func Error(msg string, c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"error": msg,
	})
}
