package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	type Person struct {
		Name  string   `json:"name"`
		Age   int      `json:"age"`
		Itens []string `json:"itens"`
	}

	app := fiber.New()

	app.Get("/structToJSON", func(c *fiber.Ctx) error {
		person := Person{}
		person.Name = "Rafael"
		person.Age = 25
		person.Itens = append(person.Itens, "phone")
		person.Itens = append(person.Itens, "mouse")

		return c.Status(fiber.StatusOK).JSON(person)
	})

	app.Post("/bodyToStruct", func(c *fiber.Ctx) error {
		person := Person{}
		err := c.BodyParser(&person)
		if err != nil {
			log.Fatal(err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(person)
	})

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err.Error())
	}

}
