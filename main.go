package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Membuat instance Fiber
	app := fiber.New()

	// Membuat route untuk halaman beranda
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world!")
	})

	// Menjalankan server di port 3000
	log.Fatal(app.Listen(":3000"))
}
