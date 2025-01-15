package database

import (
	"fmt"
	"log"
	"os"

	"github.com/ehsan-ashik/go-job-tracker-api/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var migrated bool = false

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Enable logging mode
		DryRun: false,                               // DryRun generate sql without execute
	})
	if err != nil {
		log.Fatal("Could not connect to Databse")
		log.Fatal(err.Error())
		os.Exit(2)
	}
	DB = db
	log.Println("Connected to Database at Port: ", os.Getenv("DB_PORT"))

	if !migrated {
		if err := migrate(); err != nil {
			log.Fatal("Could not migrate the database")
			log.Fatal(err.Error())
			os.Exit(2)
		}
		log.Println("Database Migrated Successfully.")
		migrated = true
	}
}

func migrate() error {
	return DB.AutoMigrate(&model.Company{}, &model.JobDescription{}, &model.JobCategory{}, &model.Job{})
}
