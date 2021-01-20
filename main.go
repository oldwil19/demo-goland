package main

import (
    //"log"
	//"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/qinains/fastergoding"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupRoutes(app *fiber.App) {
	// give response when at /
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success":  true,
			"message": "You are at the endpoint 😉",
		})
	})

/*	// api group
	api := app.Group("/api")

	// give response when at /api
	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the api endpoint 😉",
		})
	})

	// send todos route group to TodoRoutes of routes package
	routes.TodoRoute(api.Group("/todos"))*/
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

/*
func getPerson(c *fiber.Ctx) {
}
func createPerson(c *fiber.Ctx) {
}
func updatePerson(c *fiber.Ctx) {
}
func deletePerson(c *fiber.Ctx) {
}*/