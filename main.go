package main

import (
	_ "strongpasswordapi/docs"
	"strongpasswordapi/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

// @title API Gerador de Senha Forte
// @version 1.0
// @description Um gerador de senha aleatória segura.
// @host strong-password-api.onrender.com/
// @schemes https
func main() {

	app := fiber.New()

	app.Use(cors.New())

	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/password", handlers.PasswordHandler)
	app.Listen(":8080")
}
