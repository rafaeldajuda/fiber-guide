package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// CONFIG APP
	app := fiber.New()

	// ROUTE
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("server ok")
	})

	// START SERVER
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err.Error())
	}

}
