package main

import (
	"fiber-project/database"
	"fiber-project/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true, //Observar pois estava como true e não estava funcionando	
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept, Accept-Language, Content-Length",	
	}))
	routes.Setup(app)
	app.Listen(":3000")
	
}

func Home(c *fiber.Ctx) error {
	return c.SendString("Hello world!")
}

