package internal

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	// Connect to postrgres
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=postgres dbname=read_db port=5434 sslmode=disable TimeZone=Asia/Shanghai"))
	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Connection Opened to Database")

	return db
}
