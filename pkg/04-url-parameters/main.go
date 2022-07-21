package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	// WITHOUT PARAMETER
	app.Get("/parameter", func(c *fiber.Ctx) error {
		return c.SendString("route without parameter")
	})

	// WITH PARAMETER
	app.Get("/parameter/:item", func(c *fiber.Ctx) error {
		item := c.Params("item")
		return c.SendString(item)
	})

	// OPTIONAL PARAMETER
	app.Get("/optionalParameter/:item?", func(c *fiber.Ctx) error {
		item := c.Params("item")
		return c.SendString(item)
	})

	// ANY ROUTE (greedy)
	app.Get("/anyParameter/*", func(c *fiber.Ctx) error {
		item := c.Params("*")
		return c.SendString(item)
	})

	// OTHER
	app.Get("/parameterColor/color::color", func(c *fiber.Ctx) error {
		item := c.Params("color")
		return c.SendString(item)
	})

	// OTHER
	app.Get("/int/:number", func(c *fiber.Ctx) error {
		number, err := c.ParamsInt("number", 0)
		if err != nil {
			log.Fatal(err.Error())
		}

		return c.SendString(strconv.Itoa(number))
	})

	// START SERVER
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err.Error())
	}

}
