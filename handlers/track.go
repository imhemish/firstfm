package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func TrackHandler(c *fiber.Ctx) error {
	// Set by middleware
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	uid := claims["uid"]
	println(uid)

	return nil

	/* track, err := strconv.Atoi(c.Params("track"))
	if track == 0 || err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{Error: "Track empty"})
	}
	trackDetails, err := repository.GetTrackDetails(uint(track), uint(uid))
	if err != nil {
		return c.Status(404).JSON(common.ErrorResponse{
			Error: err.Error(),
		})
	}
	return c.JSON(trackDetails) */
}
