package handlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/imhemish/firstradio/common"
	"github.com/imhemish/firstradio/repository"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Signup(c *fiber.Ctx) error {
	signupRequest := SignupRequest{}
	err := c.BodyParser(&signupRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if signupRequest.Email == "" || signupRequest.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{
			Error: "Not all fields populated",
		})
	}

	if !strings.Contains(signupRequest.Email, "@") {
		return c.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{
			Error: "Email not valid",
		})
	}

	if len(signupRequest.Password) < 8 {

		return c.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{
			Error: "Password not valid",
		})
	}

	id := uuid.New().String()
	hashedpwd, _ := bcrypt.GenerateFromPassword([]byte(signupRequest.Password), 4)
	user := repository.User{UserID: id, Email: signupRequest.Email, Hashedpwd: string(hashedpwd)}
	error := repository.SignupUser(user)
	if error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{
			Error: error.Error(),
		})
	}
	token, err := common.GenerateToken(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.ErrorResponse{
			Error: err.Error(),
		})
	}

	c.JSON(fiber.Map{
		"uid":   user.UserID,
		"token": token,
	})
	return nil
}

func Login(c *fiber.Ctx) error {
	loginRequest := LoginRequest{}
	err := c.BodyParser(&loginRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{
			Error: err.Error(),
		})
	}

	user, err := repository.FindByCredentials(loginRequest.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{
			Error: err.Error(),
		})
	}

	if user.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{
			Error: "Email does not exist",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Hashedpwd), []byte(loginRequest.Password))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.ErrorResponse{
			Error: "Wrong password",
		})
	}

	token, err := common.GenerateToken(user)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(common.ErrorResponse{
			Error: err.Error(),
		})
	}

	// Return the token
	c.JSON(fiber.Map{"token": token})
	return nil
}
