package main

import (
	"Ders/egitim"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("xxx")

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(egitim.PublicVeriable)
	})

	app.Listen(":3000")
}
