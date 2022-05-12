package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	// GROUP 1
	v1 := app.Group("/v1")
	v1.Get("/group", func(c *fiber.Ctx) error {
		return c.SendString("group 1")
	})

	// GROUP 2
	v2 := app.Group("/v2")
	v2.Get("/group", func(c *fiber.Ctx) error {
		return c.SendString("group 2")
	})

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err.Error())
	}

}
