package main

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"time"
)

func main() {
	fmt.Println("Hello World")

	app := fiber.New(
		fiber.Config{
			IdleTimeout: time.Second * 5,
		})

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":3000")
}
