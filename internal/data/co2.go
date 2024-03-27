package data

import (
	"database/sql"
	"time"

	"github.com/FMinister/co2monitor-api/internal/validator"
)

type Co2Model struct {
	DB *sql.DB
}

type Co2Data struct {
	ID         int       `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	LocationID int64     `json:"location_id"`
	Co2        int       `json:"co2"`
	Temp       float64   `json:"temp"`
	Humidity   int       `json:"humidity"`
}

func ValidateCo2Data(v *validator.Validator, co2Data *Co2Data) {
	v.Check(co2Data.LocationID > 0, "location_id", "must greater than zero")

	v.Check(co2Data.Co2 >= 0, "co2", "must be greater than or equal to zero ppm")
	v.Check(co2Data.Co2 < 5000, "co2", "must be less than 5000 ppm")

	v.Check(co2Data.Temp >= -100, "temp", "must be greater than or equal to -100 degrees Celsius")
	v.Check(co2Data.Temp <= 100, "temp", "must be less than or equal to 100 degrees Celsius")

	v.Check(co2Data.Humidity >= 0, "humidity", "must be greater than or equal to zero percent")
	v.Check(co2Data.Humidity <= 100, "humidity", "must be less than or equal to 100 percent")
}
