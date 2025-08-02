package handlers

import (
	"strongpasswordapi/services"

	"github.com/gofiber/fiber/v2"
)

// @Summary Gera uma senha forte
// @Description Retorna uma senha aleatória com letras, números e símbolos
// @Tags Senha
// @Produce json
// @Success 200 {object} map[string]string
// @Router /password [get]
func PasswordHandler(c *fiber.Ctx) error {
	password, err := services.GeneratePassword()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"password": password,
	})
}
