package main

import (
	"log"

	"eagle-backend-dashboard/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Membuat koneksi dari config/database.go
	db1, err := config.NewDatabaseConfig()
	if err != nil {
		// log with comment
		log.Fatalf("Failed to load database config: %v", err)
	}

	// Connect ke database
	_, err = config.Connect(db1)
	if err != nil {
		// log with comment
		log.Fatalf(err.Error())
	}

	// Membuat instance Fiber
	app := fiber.New()

	// Membuat route untuk halaman beranda
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world!")
	})

	// Menjalankan server di port 3000
	log.Fatal(app.Listen(":3000"))
}
