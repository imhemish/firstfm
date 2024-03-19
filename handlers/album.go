package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/imhemish/firstradio/common"
	"github.com/imhemish/firstradio/repository"
)

func AlbumHandler(c *fiber.Ctx) error {
	album, err := strconv.Atoi(c.Params("album"))
	if album == 0 || err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{Error: "Album empty"})
	}
	albumDetails, err := repository.GetAlbumDetails(uint(album))
	if err != nil {
		return c.Status(404).JSON(common.ErrorResponse{
			Error: err.Error(),
		})
	}
	return c.JSON(albumDetails)

}
