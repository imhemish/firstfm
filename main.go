package main

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/imhemish/firstradio/common"
	"github.com/imhemish/firstradio/handlers"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	app := fiber.New()

	v1 := app.Group("/api/v1/")

	v1.Post("/signup", handlers.Signup)
	v1.Post("/login", handlers.Login)

	protected := v1.Group("/protected")

	protected.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte(common.Secret),
		},
	}))

	protected.Put("/profile", handlers.PutProfile)
	protected.Put("/password", handlers.PutPassword)
	protected.Get("/artist/:artist", handlers.ArtistHandler)
	protected.Get("/album/:album", handlers.AlbumHandler)
	protected.Get("/track/:track", handlers.TrackHandler)

	app.Listen(":8080")
}
