package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/get", func(c *fiber.Ctx) error {
		return c.SendString("GET route")
	})

	app.Post("/post", func(c *fiber.Ctx) error {
		body := string(c.Body())
		return c.SendString(body)
	})

	app.Put("/put", func(c *fiber.Ctx) error {
		body := string(c.Body())
		return c.SendString(body)
	})

	app.Patch("/patch", func(c *fiber.Ctx) error {
		body := string(c.Body())
		return c.SendString(body)
	})

	app.Delete("/delete", func(c *fiber.Ctx) error {
		return c.SendString("DELETE route")
	})

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err.Error())
	}

}
