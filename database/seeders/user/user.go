package main

import (
	"eagle-backend-dashboard/config"
	"eagle-backend-dashboard/entity"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Membuat koneksi dari config/database.go
	dbConfig, err := config.NewDatabaseConfig()
	if err != nil {
		// log with comment
		log.Fatalf(err.Error())
	}

	// Connect ke database
	db, err := config.Connect(dbConfig)
	if err != nil {
		// log with comment
		log.Fatalf(err.Error())
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123!"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	menus := []entity.User{
		{
			ID:       1,
			GroupID:  1,
			Name:     "Super Admin",
			Username: "superadmin",
			Password: string(hashedPassword),
			NRP:      "1234567890",
		},
	}

	for _, menu := range menus {
		err := db.FirstOrCreate(&menu).Error
		if err != nil {
			log.Fatalf(err.Error())
		}
	}

	log.Println("Users seeded successfully")
}
