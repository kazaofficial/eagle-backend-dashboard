package main

import (
	"eagle-backend-dashboard/database/seeders/seed"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		for _, arg := range args[1:] {
			switch arg {
			case "user_group":
				seed.UserGroupSeeders()
			case "user":
				seed.UserSeeders()
			case "menu":
				seed.MenuSeeders()
			case "user_group_menu":
				seed.UserGroupMenuSeeders()
			default:
				fmt.Println("Invalid seeder = " + arg)
			}
		}
	} else {
		seed.UserGroupSeeders()
		seed.UserSeeders()
		seed.MenuSeeders()
		seed.UserGroupMenuSeeders()
	}
	fmt.Println("Seeders ran successfully")
}
