package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// Field names should start with an uppercase letter
	type Person struct {
		Name     string   `query:"name"`
		Pass     string   `query:"pass"`
		Products []string `query:"products"`
	}

	app := fiber.New()

	// QUERY STRING
	app.Get("/queryString", func(c *fiber.Ctx) error {
		item := c.Query("item")
		item2 := c.Query("item2")
		return c.SendString(item + " - " + item2)
	})

	// QUERY STRING TO STRUCT
	app.Get("/queryToStruct", func(c *fiber.Ctx) error {
		p := new(Person)

		if err := c.QueryParser(p); err != nil {
			return err
		}

		log.Println(p.Name)
		log.Println(p.Pass)
		log.Println(p.Products)

		return c.Status(fiber.StatusOK).JSON(p)
	})

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err.Error())
	}

}
