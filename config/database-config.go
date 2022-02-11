package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/shwethadia/payment/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//SetupDatabaseConnection
func SetupDatabaseConnection() *gorm.DB {

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed  to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=true&loc=Local", dbUser, dbpass, dbHost, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {

		panic("Failed to create a comnnection to your database")
	}

	db.AutoMigrate(&entity.Account{}, &entity.User{}, &entity.Transaction{})
	return db

}

//CloseDatabaseConnection
func CloseDatabaseConnection(db *gorm.DB) {

	dbSQL, err := db.DB()
	if err != nil {

		panic("Failed to close connection from database")
	}

	dbSQL.Close()

}
