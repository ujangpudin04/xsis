package database

import (
	"fmt"
	"xsis/models"
	"xsis/pkg/mysql"
)

// Automatic Migration if Running App
func RunMigration() {
	mysql.DB.AutoMigrate(&models.Film{})
	// mysql.DB.AutoMigrate(&models.User{})
	// mysql.DB.AutoMigrate(&models.Profile{})

	//   if err != nil {
	//     fmt.Println(err)
	//     panic("Migration Failed")
	//   }

	fmt.Println("Migration Success")
}
