package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// Custom config
	fiberConfig := fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Custom Config Fiber",
		AppName:       "Test App v1.0.1",
	}

	app := fiber.New(fiberConfig)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("server ok")
	})

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err.Error())
	}

}
