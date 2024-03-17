package seed

import (
	"eagle-backend-dashboard/config"
	"eagle-backend-dashboard/entity"
	"log"
)

func UserGroupMenuSeeders() {
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

	userGroupMenus := []entity.UserGroupMenu{
		{
			ID:          1,
			UserGroupID: 1,
			MenuID:      1,
		},
		{
			ID:          2,
			UserGroupID: 1,
			MenuID:      2,
		},
		{
			ID:          3,
			UserGroupID: 1,
			MenuID:      3,
		},
		{
			ID:          4,
			UserGroupID: 1,
			MenuID:      4,
		},
		{
			ID:          5,
			UserGroupID: 1,
			MenuID:      5,
		},
		{
			ID:          6,
			UserGroupID: 1,
			MenuID:      6,
		},
		{
			ID:          7,
			UserGroupID: 1,
			MenuID:      7,
		},
		{
			ID:          8,
			UserGroupID: 1,
			MenuID:      8,
		},
		{
			ID:          9,
			UserGroupID: 1,
			MenuID:      9,
		},
		{
			ID:          10,
			UserGroupID: 1,
			MenuID:      10,
		},
		{
			ID:          11,
			UserGroupID: 1,
			MenuID:      11,
		},
		{
			ID:          12,
			UserGroupID: 1,
			MenuID:      12,
		},
		{
			ID:          13,
			UserGroupID: 1,
			MenuID:      13,
		},
		{
			ID:          14,
			UserGroupID: 1,
			MenuID:      14,
		},
		{
			ID:          15,
			UserGroupID: 1,
			MenuID:      15,
		},
		{
			ID:          16,
			UserGroupID: 1,
			MenuID:      16,
		},
		{
			ID:          17,
			UserGroupID: 1,
			MenuID:      17,
		},
		{
			ID:          18,
			UserGroupID: 1,
			MenuID:      18,
		},
		{
			ID:          19,
			UserGroupID: 1,
			MenuID:      19,
		},
		{
			ID:          20,
			UserGroupID: 1,
			MenuID:      20,
		},
		{
			ID:          21,
			UserGroupID: 1,
			MenuID:      21,
		},
		{
			ID:          22,
			UserGroupID: 1,
			MenuID:      22,
		},
		{
			ID:          23,
			UserGroupID: 1,
			MenuID:      23,
		},
		{
			ID:          24,
			UserGroupID: 1,
			MenuID:      24,
		},
		{
			ID:          25,
			UserGroupID: 1,
			MenuID:      25,
		},
		{
			ID:          26,
			UserGroupID: 1,
			MenuID:      26,
		},
		{
			ID:          27,
			UserGroupID: 1,
			MenuID:      27,
		},
		{
			ID:          28,
			UserGroupID: 1,
			MenuID:      28,
		},
		{
			ID:          29,
			UserGroupID: 1,
			MenuID:      29,
		},
		{
			ID:          30,
			UserGroupID: 1,
			MenuID:      30,
		},
		{
			ID:          31,
			UserGroupID: 1,
			MenuID:      31,
		},
		{
			ID:          32,
			UserGroupID: 1,
			MenuID:      32,
		},
		{
			ID:          33,
			UserGroupID: 1,
			MenuID:      33,
		},
		{
			ID:          34,
			UserGroupID: 1,
			MenuID:      34,
		},
		{
			ID:          35,
			UserGroupID: 1,
			MenuID:      35,
		},
		{
			ID:          36,
			UserGroupID: 1,
			MenuID:      36,
		},
		{
			ID:          37,
			UserGroupID: 1,
			MenuID:      37,
		},
	}

	for _, userGroupMenu := range userGroupMenus {
		err := db.FirstOrCreate(&userGroupMenu).Error
		if err != nil {
			log.Fatalf(err.Error())
		}
	}

	for _, userGroupMenu := range userGroupMenus {
		err := db.FirstOrCreate(&userGroupMenu).Error
		if err != nil {
			log.Fatalf(err.Error())
		}
	}

	log.Println("User groups seeded successfully")
}
