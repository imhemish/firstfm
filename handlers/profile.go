package handlers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/imhemish/firstradio/repository"
)

type ProfileRequest struct {
	About    string `json:"about"`
	Name     string `json:"name"`
	Language string `json:"language"`
	Picture  string `json:"picture"`
}

func PutProfile(c *fiber.Ctx) error {
	profileRequest := ProfileRequest{}
	err := c.BodyParser(&profileRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user := repository.User{
		About:    sql.NullString{String: profileRequest.About, Valid: true},
		Language: sql.NullString{String: profileRequest.Language, Valid: true},
		Name:     sql.NullString{String: profileRequest.Name, Valid: true},
		Picture:  sql.NullString{String: profileRequest.Picture, Valid: true},
	}
	user.UserID = c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["uid"].(string)
	err = repository.UpdateProfile(user)
	return err
}
