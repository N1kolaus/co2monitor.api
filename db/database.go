package db

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
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

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: false, // disables implicit prepared statement usage
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "co2monitor.", // schema name
			SingularTable: false,
		},
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	dbSchema := os.Getenv("POSTGRES_DB")
	createSchemaCommand := fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s;", dbSchema)
	result := db.Exec(createSchemaCommand)
	if result.Error != nil {
		log.Fatal("Failed to create schema. \n", result.Error)
	}

	log.Info("Connected to database.")

	DB = DbInstance{
		Db: db,
	}
}

func GetDB() *gorm.DB {
	return DB.Db
}
