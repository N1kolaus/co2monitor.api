package initializers

import (
	"github.com/fminister/co2monitor.api/db"
	"github.com/fminister/co2monitor.api/models"
)

func SyncDatabase() {
	db := db.GetDB()
	db.AutoMigrate(
		&models.Co2Data{},
		&models.Location{},
	)
}
