package controllers

import (
	"fiber-project/models"

	"github.com/gofiber/fiber/v2"
) 

func Register(c *fiber.Ctx) error {
	user := models.User{
		FirstName: "Toti",
	}

	user.LastName = "Cavalcante"
	return c.JSON(user)

}