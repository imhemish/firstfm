package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/imhemish/firstradio/common"
	"github.com/imhemish/firstradio/repository"
	"golang.org/x/crypto/bcrypt"
)

type PasswordRequest struct {
	OldPassword string `json:"old"`
	NewPassword string `json:"new"`
}

func PutPassword(c *fiber.Ctx) error {
	passwordRequest := PasswordRequest{}
	err := c.BodyParser(&passwordRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	id := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["uid"].(string)

	if len(passwordRequest.NewPassword) < 8 {
		return c.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{
			Error: "too short new password",
		})
	}
	oldhashedpwd, err := repository.GetHashedPassword(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.ErrorResponse{
			Error: err.Error(),
		})
	}
	err = bcrypt.CompareHashAndPassword([]byte(oldhashedpwd), []byte(passwordRequest.OldPassword))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{
			Error: "old password does not match",
		})
	}
	hashedpwd, _ := bcrypt.GenerateFromPassword([]byte(passwordRequest.NewPassword), 4)
	err = repository.UpdatePassword(id, string(hashedpwd))
	return err
}
