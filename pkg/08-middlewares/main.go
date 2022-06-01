package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New()

	// Default middleware config
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("server ok")
	})

	// Custom middleware
	app.Use(customMiddlware)

	app.Post("/", func(c *fiber.Ctx) error {
		return c.SendString(string(c.Body()))
	})

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err.Error())
	}

}

// customMiddlware - check if the body request is not empty
func customMiddlware(c *fiber.Ctx) error {
	body := string(c.Body())
	if len(body) == 0 {
		c.Set("Content-Type", "application/json")
		return c.Status(fiber.StatusBadRequest).SendString(`{"msg": "invalid body"}`)
	}
	return c.Next()
}
