package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"learn-go-auth/models"
)

var DB *gorm.DB

// ConnectDatabase initializes the MySQL database connection
func ConnectDatabase() {
	var err error

	// MySQL connection string format: username:password@protocol(address)/dbname?param=value
	dsn := "root:@tcp(127.0.0.1:3306)/learn_go_auth?charset=utf8mb4&parseTime=True&loc=Local"

	// Open the MySQL database connection
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Auto migrate the User model
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("MySQL database connection successfully established")
}
