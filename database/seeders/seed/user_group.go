package seed

import (
	"eagle-backend-dashboard/config"
	"eagle-backend-dashboard/entity"
	"log"
)

func UserGroupSeeders() {
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

	userGroups := []entity.UserGroup{
		{
			ID:          1,
			Name:        "Admin",
			Description: "Admin user group",
		},
		{
			ID:          2,
			Name:        "User",
			Description: "User user group",
		},
	}

	for _, userGroup := range userGroups {
		err := db.FirstOrCreate(&userGroup).Error
		if err != nil {
			log.Fatalf(err.Error())
		}
	}

	log.Println("User groups seeded successfully")
}
