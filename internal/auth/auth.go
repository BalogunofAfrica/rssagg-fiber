package auth

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// GetAPIKey - extracts api key from request header
// Example:
// Authorization: Apikey (insert key here)
func GetAPIKey(c *fiber.Ctx) (string, error) {
	val := c.Get("Authorization")
	if val == "" {
		return "", errors.New(("no authentication info found"))
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New(("malformed authentication header"))
	}

	return vals[1], nil
}
