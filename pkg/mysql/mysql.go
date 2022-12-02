package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connection Database
func DatabaseInit() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/xsis?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// Get `host`, `user`, `password`, `database name`, and `port` from env here ...

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to Database")
}
