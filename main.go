package main

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"

	"eagle-backend-dashboard/config"
	"eagle-backend-dashboard/controller"
	"eagle-backend-dashboard/dto"
	"eagle-backend-dashboard/middleware"
	"eagle-backend-dashboard/repository"
	"eagle-backend-dashboard/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	app.Use(cors.New())

	// Middleware
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	app.Use(func(c *fiber.Ctx) error {
		// Catch panics
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}

				// Get the stack trace
				stackTrace := debug.Stack()

				// Log the error along with the stack trace
				log.Printf("%v\n%s", err, stackTrace)

				// Send a response to the client
				c.Status(500).JSON(dto.ErrorResponse{
					StatusCode: 500,
					Message:    "Internal Server Error",
					Error:      err.Error(),
				})
			}
		}()

		return c.Next()
	})

	// initialize repositories
	userGroupRepository := repository.NewUserGroupRepository(db)
	menuRepository := repository.NewMenuRepository(db)
	userRepository := repository.NewUserRepository(db)
	userGroupMenuRepository := repository.NewUserGroupMenuRepository(db)
	daftarProsesPenarikanDataRepository := repository.NewDaftarProsesPenarikanDataRepository(db)

	// initialize services
	userGroupService := service.NewUserGroupService(userGroupRepository, menuRepository)
	menuService := service.NewMenuService(menuRepository)
	authService := service.NewAuthService(userRepository)
	userService := service.NewUserService(userRepository)
	userGroupMenuService := service.NewUserGroupMenuService(userGroupMenuRepository, menuRepository)
	supersetService := service.NewSupersetService()
	olahDataService := service.NewOlahDataService(daftarProsesPenarikanDataRepository)

	// add middleware
	app.Use(middleware.AuthenticationMiddleware())

	// initialize controllers group for api/v1
	apiv1 := app.Group("/api/v1")
	controller.NewUserGroupRoutes(apiv1, userGroupService)
	controller.NewMenuRoutes(apiv1, menuService)
	controller.NewAuthRoutes(apiv1, authService)
	controller.NewUserRoutes(apiv1, userService)
	controller.NewUserGroupMenuRoutes(apiv1, userGroupMenuService)
	controller.NewSupersetRoutes(apiv1, supersetService)
	controller.NewOlahDataRoutes(apiv1, olahDataService)

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
