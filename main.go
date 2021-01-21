package main

import (
    //"log"
	//"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/qinains/fastergoding"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"demo/controllers"
)

func setupRoutes(app *fiber.App) {
	// give response when at /
	/*app.Post("/topsecret", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success":  true,
			"message": "You are at the endpoint 😉",
		})
	})*/
	app.Post("/topsecret/", controllers.TopSecret)
	app.Post("/topsecret_split/:satelliteName", controllers.TopSecretByOne)
}



func main() {
	fastergoding.Run() // hot reload
    app := fiber.New()

	app.Use(logger.New())

	// setup routes
	setupRoutes(app)

	// Listen on server 8000 and catch error if any
	err := app.Listen(":8088")

	// handle error
	if err != nil {
		panic(err)
	} 
}
