package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

func ConnectToDb() {
	dsn := os.Getenv("DATABASE_URL_DEV")

	if os.Getenv("APP_ENV") != "development" {
		dsn = os.Getenv("DATABASE_URL")
	}

	log.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "co2monitor", // schema name
			SingularTable: false,
		},
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	// db.AutoMigrate(&models.Book{})

	DB = DbInstance{
		Db: db,
	}
}
