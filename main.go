package main

import (
	"log"
	"os"

	"eagle-backend-dashboard/config"
	"eagle-backend-dashboard/controller"
	"eagle-backend-dashboard/dto"
	"eagle-backend-dashboard/repository"
	"eagle-backend-dashboard/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	// Membuat koneksi dari config/database.go
	db1, err := config.NewDatabaseConfig()
	if err != nil {
		// log with comment
		log.Fatalf("Failed to load database config: %v", err)
	}

	// Connect ke database
	db, err := config.Connect(db1)
	if err != nil {
		// log with comment
		log.Fatalf(err.Error())
	}

	// Membuat instance Fiber
	app := fiber.New()

	// cors handler
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		return c.Next()
	})

	// Middleware
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	// initialize repositories
	userGroupRepository := repository.NewUserGroupRepository(db)
	menuRepository := repository.NewMenuRepository(db)

	// initialize services
	userGroupService := service.NewUserGroupService(userGroupRepository)
	menuService := service.NewMenuService(menuRepository)

	// initialize controllers group for api/v1
	apiv1 := app.Group("/api/v1")
	controller.NewUserGroupRoutes(apiv1, userGroupService)
	controller.NewMenuRoutes(apiv1, menuService)

	// Add a middleware for handling not found errors
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(dto.ErrorResponse{
			StatusCode: 404,
			Message:    "Not Found",
			Error:      "Endpoint not found",
		})
	})

	// Menjalankan server di port 3000
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}
	err = app.Listen(":" + port)

	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
