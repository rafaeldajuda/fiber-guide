package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	// GET REQUEST HEADERS
	app.Get("/requestHeaders", func(c *fiber.Ctx) error {
		headers := c.GetReqHeaders()
		return c.Status(fiber.StatusOK).JSON(headers)
	})

	// GET SPECIFIC REQUEST HEADER
	app.Get("/requestSpecificHeader", func(c *fiber.Ctx) error {
		headers := c.GetReqHeaders()
		host := headers["Host"]
		return c.Status(fiber.StatusOK).SendString(host)
	})

	// GET RESPONSE HEADER
	app.Get("/responseHeaders", func(c *fiber.Ctx) error {
		headers := c.GetRespHeaders()
		return c.Status(fiber.StatusOK).JSON(headers)
	})

	// SET RESPONSE HEADER
	app.Get("/setResponseHeader", func(c *fiber.Ctx) error {
		c.Response().Header.Add("data", time.Now().Format(time.RFC3339))
		c.Response().Header.Add("cod", "123")
		headers := c.GetRespHeaders()
		return c.Status(fiber.StatusOK).JSON(headers)
	})

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err.Error())
	}

}
