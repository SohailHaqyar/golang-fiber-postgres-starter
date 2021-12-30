package main

import (
	"fmt"
	"os"

	"github.com/SohailHaqyar/wait-so/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDatabase() {
	var err error
	URI := os.Getenv("DATABASE_URL")

	database.DatabaseConfig, err = gorm.Open(postgres.Open(URI), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Error while connecting to the database")
	}
	fmt.Println("Connected to the database successfully")
	// database.DatabaseConfig.AutoMigrate()
}
func main() {
	initDatabase()
	fmt.Println("Build away")
}
