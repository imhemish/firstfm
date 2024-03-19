package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type ArtistResponse struct {
	Error       string `json:"error"`
	GlobalPlays int    `json:"global_plays"`
	UserPlays   int    `json:"user_plays"`
	Likes       int    `json:"likes"`
}

func ArtistHandler(c *fiber.Ctx) error {
	artist := c.Params("artist")
	if artist == "" {
		c.JSON(ArtistResponse{"Artist is empty", 0, 0, 0})
	}
	return nil
}
